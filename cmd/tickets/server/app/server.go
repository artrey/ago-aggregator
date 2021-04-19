package app

import (
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type Server struct {
	ticketsV1Pb.UnimplementedTicketsServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Search(request *ticketsV1Pb.SearchRequest, server ticketsV1Pb.TicketsService_SearchServer) error {
	log.Println(request)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		if err := server.Send(&ticketsV1Pb.SearchResponse{
			Id:            int64(i + 1),
			DepartureTime: timestamppb.New(time.Now()),
			TravelTime:    &timestamppb.Timestamp{Seconds: int64(time.Hour * 2 / time.Second)},
			Cost:          20_000_00,
		}); err != nil {
			log.Println(err)
			return err
		}
	}
	log.Println("Successfully finished")
	return nil
}
