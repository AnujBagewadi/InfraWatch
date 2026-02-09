package main

import (
	"log"
	"os"

	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
	"github.com/AnujBagewadi/InfraWatch/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using system environment variables")
	}

	repository.InitMongo()

	app := fiber.New(fiber.Config{
		AppName: "Real-Time Monitoring Dashboard v1",
	})

	routes.RegisterAPIRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
