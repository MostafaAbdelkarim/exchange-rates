package route

import (
	"ExchangeRates/controller"
	middleware "ExchangeRates/middleware"

	"github.com/gofiber/fiber/v2"
)

func EthereumBalanceRoutes(app *fiber.App) {

	api := app.Group("/api/v1/balance")

	api.Get("/:address", middleware.ValidateAddress, controller.GetBalanceUsingAddress)

}
