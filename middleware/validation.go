package middleware

import (
	errorHandler "ExchangeRates/error_handling"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ValidateCryptoCurrency(c *fiber.Ctx) error {

	cryptocurrerncy := c.Params("cryptocurrerncy")

	if len(cryptocurrerncy) > 15 {
		return errorHandler.HandleErrorString(c, http.StatusBadRequest, "Crypto Currency param should not be longer than 15 characters")
	}

	return c.Next()
}

func ValidateFiatCurrency(c *fiber.Ctx) error {

	fiat := c.Params("fiat")

	if len(fiat) > 15 {
		return errorHandler.HandleErrorString(c, http.StatusBadRequest, "Fiat Currency param should not be longer than 15 characters")
	}

	return c.Next()
}

func ValidateAddress(c *fiber.Ctx) error {
	address := c.Params("address")

	if len(address) > 15 {
		return errorHandler.HandleErrorString(c, http.StatusBadRequest, "Address param should not be longer than 15 characters")
	}

	return c.Next()
}
