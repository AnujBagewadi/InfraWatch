package routes

import (
	"github.com/AnujBagewadi/InfraWatch/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterLogRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/logs/:apiId", handler.GetLogsByAPI)
}
