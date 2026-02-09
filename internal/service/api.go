package service

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/models"
	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
)

func CreateAPI(ctx context.Context, name, url string, threshold int) error {
	api := models.API{
		Name:        name,
		URL:         url,
		ThresholdMS: threshold,
		CreatedAt:   time.Now(),
	}

	return repository.InsertAPI(ctx, api)
}

func ListAPIs(ctx context.Context) ([]models.API, error) {
	return repository.GetAllAPIs(ctx)
}
