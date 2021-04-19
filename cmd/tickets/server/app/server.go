package app

import (
	ticketsV1Pb "github.com/artrey/ago-aggregator/pkg/api/proto/v1"
	"github.com/artrey/ago-aggregator/pkg/search"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sync"
)

type Server struct {
	ticketsV1Pb.UnimplementedTicketsServiceServer
	backends []search.Backend
}

func New(backends ...search.Backend) *Server {
	return &Server{
		backends: backends,
	}
}

func (s *Server) Search(request *ticketsV1Pb.SearchRequest, server ticketsV1Pb.TicketsService_SearchServer) error {
	log.Println(request)

	sendCh := make(chan search.Journey, len(s.backends))
	wg := sync.WaitGroup{}

	for _, backend := range s.backends {
		channel, err := backend.Search(server.Context(), request.Date.AsTime(), request.CityFrom, request.CityTo)
		if err != nil {
			log.Println(err)
			continue
		}

		wg.Add(1)
		go func(ch <-chan search.Journey) {
			defer wg.Done()
			for journey := range ch {
				sendCh <- journey
			}
		}(channel)
	}

	go func() {
		for journey := range sendCh {
			err := server.Send(&ticketsV1Pb.SearchResponse{
				Id:            journey.Id,
				DepartureTime: timestamppb.New(journey.Start),
				TravelTime:    &timestamppb.Timestamp{Seconds: journey.Duration},
				Cost:          journey.Cost,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}()

	wg.Wait()
	close(sendCh)

	log.Println("Finished: ", request)
	return nil
}
