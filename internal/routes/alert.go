package routes

import (
	"github.com/AnujBagewadi/InfraWatch/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlertRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/alerts", handler.GetAlerts)
}
