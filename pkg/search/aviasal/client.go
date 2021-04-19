package aviasal

import (
	"github.com/artrey/ago-aggregator/pkg/search"
	"math/rand"
	"time"
)

type Aviasal struct {
	*search.BaseClient
}

func New(bufferSize int) *Aviasal {
	data := []search.Journey{
		{
			Id:       rand.Int63n(100_000_000),
			From:     "Moscow",
			To:       "Kazan",
			Start:    time.Date(2021, time.April, 30, 12, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour * 2),
			Cost:     7_000_00,
		},
		{
			Id:       rand.Int63n(100_000_000),
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 15, 30, 0, 0, time.UTC),
			Duration: int64(time.Hour * 2),
			Cost:     6_500_00,
		},
		{
			Id:       rand.Int63n(100_000_000),
			From:     "Moscow",
			To:       "Saint Petersburg",
			Start:    time.Date(2021, time.April, 30, 22, 0, 0, 0, time.UTC),
			Duration: int64(time.Hour * 2),
			Cost:     9_000_00,
		},
	}

	return &Aviasal{search.NewBaseClient(bufferSize, data)}
}
