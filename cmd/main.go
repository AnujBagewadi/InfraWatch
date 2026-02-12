package main

import (
	"log"
	"os"

	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
	"github.com/AnujBagewadi/InfraWatch/internal/routes"
	"github.com/AnujBagewadi/InfraWatch/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "*",
	}))
	routes.RegisterAPIRoutes(app)
	routes.RegisterLogRoutes(app)
	routes.RegisterDashboardRoutes(app)
	routes.RegisterAlertRoutes(app)

	// Start monitoring engine
	service.StartMonitorEngine()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
