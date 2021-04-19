package search

import (
	"context"
	"time"
)

type Journey struct {
	Id       int64
	From     string
	To       string
	Start    time.Time
	Duration int64
	Cost     uint64
}

type Backend interface {
	Search(ctx context.Context, date time.Time, from, to string) (<-chan Journey, error)
}
