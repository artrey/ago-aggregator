package skyner

import (
	"github.com/artrey/ago-aggregator/pkg/search"
	"math/rand"
	"time"
)

type Undex struct {
	*search.BaseClient
}

func New(bufferSize int) *Undex {
	data := []search.Journey{
		{
			Id:       rand.Int63n(100_000_000) + 200_000_000,
			From:     "Moscow",
			To:       "Kazan",
			Start:    time.Date(2021, time.April, 30, 21, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour * 3),
			Cost:     8_000_00,
		},
		{
			Id:       rand.Int63n(100_000_000) + 200_000_000,
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 16, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour * 3),
			Cost:     5_500_00,
		},
		{
			Id:       rand.Int63n(100_000_000) + 200_000_000,
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 10, 0, 0, 0, time.UTC),
			Duration: int64(time.Hour * 3),
			Cost:     5_000_00,
		},
	}

	return &Undex{search.NewBaseClient(bufferSize, data)}
}
