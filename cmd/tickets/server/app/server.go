package app

import (
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"log"
)

type Server struct {
	ticketsV1Pb.UnimplementedTicketsServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Search(request *ticketsV1Pb.SearchRequest, server ticketsV1Pb.TicketsService_SearchServer) error {
	log.Println(request)
	return nil
}
