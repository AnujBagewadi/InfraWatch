package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// 2. Get variables from .env
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Fallback port
	}

	// 3. Connect to MongoDB
	db := repository.ConnectDB(mongoURI, dbName)
	_ = db // We will use this 'db' variable for our repositories later

	// 4. Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName: "Real-Time Monitoring Dashboard v1",
	})

	// 5. Basic Health Check Route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "UP",
			"message": "Monitoring Server is Running",
		})
	})

	// 6. Start Server
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
