package service

import (
	"context"
	"time"

	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
)

func FetchLogsByAPI(ctx context.Context, apiID string) (interface{}, error) {
	return repository.GetLogsByAPI(ctx, apiID)
}

func FetchLogsByTime(ctx context.Context, from, to time.Time) (interface{}, error) {
	return repository.GetLogsByTimeRange(ctx, from, to)
}
