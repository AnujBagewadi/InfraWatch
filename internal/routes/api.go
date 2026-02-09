package routes

import (
	"github.com/AnujBagewadi/InfraWatch/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/apis", handler.AddAPI)
	api.Get("/apis", handler.GetAPIs)
}
