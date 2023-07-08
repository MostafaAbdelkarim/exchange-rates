package main

import (
	"ExchangeRates/config"
	"ExchangeRates/cron"
	route "ExchangeRates/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New(config.GetLoggerConfig()))
	app.Use(cors.New(config.GetApiConfig()))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"Hello": "Exchange Rates"})
	})

	cron.UpdateExchangeRatesEvery5Minutes()

	route.RatesRoutes(app)
	route.EthereumBalanceRoutes(app)

	app.Listen(":3000")
}
