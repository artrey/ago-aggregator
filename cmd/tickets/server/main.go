package main

import (
	"github.com/artrey/ago-aggregator/cmd/tickets/server/app"
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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
	server := app.New()

	grpcServer := grpc.NewServer()
	ticketsV1Pb.RegisterTicketsServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return grpcServer.Serve(listener)
}
