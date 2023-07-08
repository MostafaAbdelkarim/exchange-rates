package route

import (
	"ExchangeRates/controller"
	validation "ExchangeRates/middleware"

	"github.com/gofiber/fiber/v2"
)

func RatesRoutes(app *fiber.App) {

	api := app.Group("/api/v1/rates")

	api.Get("/:cryptocurrerncy/:fiat", validation.ValidateCryptoCurrency, validation.ValidateFiatCurrency, controller.GetExchangeRateGivenCryptoAndFiat)

	api.Get("/:cryptocurrerncy", validation.ValidateCryptoCurrency, controller.GetExchangeRateBetweenCryptoAndAllFiat)

	api.Get("/", controller.GetAllExchangeRatesBetweenCryptoAndFiat)

	api.Get("/history/:cryptocurrerncy/:fiat", validation.ValidateCryptoCurrency, validation.ValidateFiatCurrency, controller.GetExchangeRateHistoryBetweenCryptoAndFiat)

}
