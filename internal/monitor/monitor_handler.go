package monitor

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	res := &Monitor{
		Name:    "live-streaming",
		Version: "0.0.0",
	}
	return c.JSON(res)
}
