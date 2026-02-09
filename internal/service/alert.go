package service

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/models"
	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
)

func CreateAlert(apiID, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	alert := models.Alert{
		APIID:     apiID,
		Message:   message,
		Timestamp: time.Now(),
		Resolved:  false,
	}

	repository.InsertAlert(ctx, alert)
}

func FetchAlerts(ctx context.Context) ([]models.Alert, error) {
	return repository.GetAlerts(ctx)
}

func ShouldTriggerFailureAlert(apiID string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := repository.GetLastNLogs(ctx, apiID, 3)
	if err != nil || len(logs) < 3 {
		return false
	}

	for _, l := range logs {
		if l.Success {
			return false
		}
	}

	return true
}
