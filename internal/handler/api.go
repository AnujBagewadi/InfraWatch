package handler

import (
	"context"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AddAPI(c *fiber.Ctx) error {
	type reqBody struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		ThresholdMS int    `json:"threshold_ms"`
	}

	var body reqBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := service.CreateAPI(ctx, body.Name, body.URL, body.ThresholdMS)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "API added successfully"})
}

func GetAPIs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	apis, err := service.ListAPIs(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(apis)
}
