package error

import (
	schema "ExchangeRates/schema/base"

	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(schema.ErrorResponse{Status: "Failed", StatusCode: status, Error: err.Error()})
}

func HandleErrorString(c *fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(schema.ErrorResponse{Status: "Failed", StatusCode: status, Error: err})
}
