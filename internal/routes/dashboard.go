package routes

import (
	"github.com/AnujBagewadi/InfraWatch/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/dashboard/:apiId", handler.GetAPIDashboard)
}
