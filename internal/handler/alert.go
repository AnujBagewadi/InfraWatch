package handler

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetAlerts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	alerts, err := service.FetchAlerts(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(alerts)
}
