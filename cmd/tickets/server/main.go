package main

import (
	"github.com/artrey/ago-aggregator/cmd/tickets/server/app"
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"github.com/artrey/ago-aggregator/pkg/search/aviasal"
	"github.com/artrey/ago-aggregator/pkg/search/skyner"
	"github.com/artrey/ago-aggregator/pkg/search/undex"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = "9999"
)

func main() {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string) error {
	rand.Seed(time.Now().Unix())
	server := app.New(aviasal.New(3), undex.New(5), skyner.New(10))

	grpcServer := grpc.NewServer()
	ticketsV1Pb.RegisterTicketsServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return grpcServer.Serve(listener)
}
