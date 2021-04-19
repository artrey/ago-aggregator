package undex

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
			Id:       rand.Int63n(100_000_000) + 100_000_000,
			From:     "Moscow",
			To:       "Kazan",
			Start:    time.Date(2021, time.April, 30, 2, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour),
			Cost:     9_200_00,
		},
		{
			Id:       rand.Int63n(100_000_000) + 100_000_000,
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 6, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour),
			Cost:     8_500_00,
		},
		{
			Id:       rand.Int63n(100_000_000) + 100_000_000,
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 18, 0, 0, 0, time.UTC),
			Duration: int64(time.Hour),
			Cost:     12_000_00,
		},
	}

	return &Undex{search.NewBaseClient(bufferSize, data)}
}
