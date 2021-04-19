package main

import (
	"context"
	"errors"
	"fmt"
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type CLI struct {
}

const defaultTimeout = 10

var (
	app = kingpin.New("client", "A command-line client to tickets aggregator grpc server.")

	host = app.Flag("host", "Host of grpc server.").Default("localhost").String()
	port = app.Flag("port", "Port of grpc server.").Default("9999").String()

	search        = app.Command("search", "Search tickets.")
	searchDate    = search.Arg("date", "Departure date in format YYYY-MM-DD.").Required().String()
	searchFrom    = search.Arg("from", "Departure city.").Required().String()
	searchTo      = search.Arg("to", "Arrival city.").Required().String()
	searchTimeout = search.Arg("timeout", "Maximum search time in seconds.").Default(strconv.Itoa(defaultTimeout)).Int()
)

func (cli *CLI) Run() (err error) {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	conn, err := grpc.Dial(net.JoinHostPort(*host, *port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
		}
	}()

	client := ticketsV1Pb.NewTicketsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*searchTimeout))
	defer cancel()

	switch cmd {
	case search.FullCommand():
		date, err := time.Parse("2006-01-02", *searchDate)
		if err != nil {
			return err
		}

		stream, err := client.Search(ctx, &ticketsV1Pb.SearchRequest{
			Date:     timestamppb.New(date),
			CityFrom: *searchFrom,
			CityTo:   *searchTo,
		})
		if err != nil {
			return err
		}

		for {
			response, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return err
			}

			fmt.Printf("%+v\n", response)
		}
	}

	fmt.Println("That's all")

	return nil
}
