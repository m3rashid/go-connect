package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleControllerError(c *fiber.Ctx, err error) error {
	log.Println(err)
	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}
