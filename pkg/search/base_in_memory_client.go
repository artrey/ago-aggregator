package search

import (
	"context"
	"log"
	"math/rand"
	"time"
)

type BaseClient struct {
	bufferSize int
	data       []Journey
}

func NewBaseClient(bufferSize int, data []Journey) *BaseClient {
	return &BaseClient{
		bufferSize: bufferSize,
		data:       data,
	}
}

func (bc *BaseClient) Search(ctx context.Context, date time.Time, from, to string) (<-chan Journey, error) {
	ch := make(chan Journey, bc.bufferSize)
	go func() {
		defer close(ch)
		for _, journey := range bc.data {
			if err := ctx.Err(); err != nil {
				log.Println(err)
				return
			}

			jYear, jMonth, jDay := journey.Start.Date()
			dYear, dMonth, dDay := date.Date()
			if jYear == dYear && jMonth == dMonth && jDay == dDay && journey.From == from && journey.To == to {
				ch <- journey
			}

			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		}
	}()
	return ch, nil
}
