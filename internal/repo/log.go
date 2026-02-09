package repository

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertLog(ctx context.Context, log models.APILog) error {
	_, err := MongoClient.
		Database("monitoring").
		Collection("logs").
		InsertOne(ctx, log)

	return err
}

func GetLogsByAPI(ctx context.Context, apiID string) ([]models.APILog, error) {
	filter := bson.M{"api_id": apiID}

	opts := options.Find().
		SetSort(bson.M{"timestamp": -1}).
		SetLimit(100)

	cursor, err := MongoClient.
		Database("monitoring").
		Collection("logs").
		Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []models.APILog
	err = cursor.All(ctx, &logs)
	return logs, err
}

func GetLogsByTimeRange(ctx context.Context, from, to time.Time) ([]models.APILog, error) {
	filter := bson.M{
		"timestamp": bson.M{
			"$gte": from,
			"$lte": to,
		},
	}

	cursor, err := MongoClient.
		Database("monitoring").
		Collection("logs").
		Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []models.APILog
	err = cursor.All(ctx, &logs)
	return logs, err
}

func GetAPISummary(ctx context.Context, apiID string) (bson.M, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"api_id": apiID}},
		{
			"$group": bson.M{
				"_id":            "$api_id",
				"avg_response":   bson.M{"$avg": "$response_time_ms"},
				"total_requests": bson.M{"$sum": 1},
				"success_count":  bson.M{"$sum": bson.M{"$cond": []interface{}{"$success", 1, 0}}},
				"failure_count":  bson.M{"$sum": bson.M{"$cond": []interface{}{"$success", 0, 1}}},
			},
		},
	}

	cursor, err := MongoClient.
		Database("monitoring").
		Collection("logs").
		Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return bson.M{}, nil
	}

	return result[0], nil
}

func GetLastNLogs(ctx context.Context, apiID string, n int64) ([]models.APILog, error) {
	filter := bson.M{"api_id": apiID}

	opts := options.Find().
		SetSort(bson.M{"timestamp": -1}).
		SetLimit(n)

	cursor, err := MongoClient.
		Database("monitoring").
		Collection("logs").
		Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []models.APILog
	err = cursor.All(ctx, &logs)
	return logs, err
}
