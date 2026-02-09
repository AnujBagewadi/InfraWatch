package repository

import (
	"context"

	"github.com/AnujBagewadi/InfraWatch/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertAlert(ctx context.Context, alert models.Alert) error {
	_, err := MongoClient.
		Database("monitoring").
		Collection("alerts").
		InsertOne(ctx, alert)

	return err
}

func GetAlerts(ctx context.Context) ([]models.Alert, error) {
	cursor, err := MongoClient.
		Database("monitoring").
		Collection("alerts").
		Find(ctx, bson.M{"resolved": false})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alerts []models.Alert
	err = cursor.All(ctx, &alerts)
	return alerts, err
}
