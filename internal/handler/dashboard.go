package handler

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetAPIDashboard(c *fiber.Ctx) error {
	apiID := c.Params("apiId")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := service.GetAPIDashboard(ctx, apiID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}
