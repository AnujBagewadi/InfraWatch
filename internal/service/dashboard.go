package service

import (
	"context"

	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
)

func GetAPIDashboard(ctx context.Context, apiID string) (interface{}, error) {
	return repository.GetAPISummary(ctx, apiID)
}
