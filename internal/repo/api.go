package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var APIColl *mongo.Collection

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	APIColl = client.Database("monitoring").Collection("apis")

	log.Println("MongoDB connected")
}

// Insert API
func InsertAPI(ctx context.Context, api models.API) error {
	_, err := APIColl.InsertOne(ctx, api)
	return err
}

// Get all APIs
func GetAllAPIs(ctx context.Context) ([]models.API, error) {
	cursor, err := APIColl.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var apis []models.API
	if err := cursor.All(ctx, &apis); err != nil {
		return nil, err
	}
	return apis, nil
}
