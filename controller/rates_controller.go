package controller

import (
	database "ExchangeRates/database"
	errorHandler "ExchangeRates/error_handling"
	exchange_mapper "ExchangeRates/mapper"
	schema "ExchangeRates/schema/base"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	cryptocompare "github.com/lucazulian/cryptocomparego"
	"go.mongodb.org/mongo-driver/mongo"
)

var exchangeCollection *mongo.Collection = database.GetCollection(database.DB, "exchange")

func GetExchangeRateGivenCryptoAndFiat(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	crypto := c.Params("cryptocurrerncy")
	fiat := c.Params("fiat")
	defer cancel()

	client := cryptocompare.NewClient(http.DefaultClient)
	prices, listResponse, err := client.Price.List(ctx, &cryptocompare.PriceRequest{Fsym: crypto, Tsyms: strings.Split(fiat, ",")})

	if err != nil || listResponse.StatusCode != 200 {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	models := exchange_mapper.ToEntityList(prices)
	interfaceSlice := make([]interface{}, len(models))
	for i, model := range models {
		interfaceSlice[i] = model
	}
	result, err := exchangeCollection.InsertMany(ctx, interfaceSlice)
	if err != nil {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	mapped := exchange_mapper.ToDtoList(prices, *result)
	response := schema.NewSuccessResponse(mapped, "Success", http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}

func GetExchangeRateBetweenCryptoAndAllFiat(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	crypto := c.Params("cryptocurrerncy")
	defer cancel()

	client := cryptocompare.NewClient(http.DefaultClient)
	prices, listResponse, err := client.Price.List(ctx, &cryptocompare.PriceRequest{Fsym: crypto, Tsyms: strings.Split("", ",")})

	if err != nil || listResponse.StatusCode != 200 {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	response := schema.NewSuccessResponse(prices, "Success", http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}

func GetAllExchangeRatesBetweenCryptoAndFiat(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := cryptocompare.NewClient(http.DefaultClient)
	prices, listResponse, err := client.Coin.List(ctx)

	if err != nil || listResponse.StatusCode != 200 {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	response := schema.NewSuccessResponse(prices, "Success", http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}

func GetExchangeRateHistoryBetweenCryptoAndFiat(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	crypto := c.Params("cryptocurrerncy")
	fiat := c.Params("fiat")
	defer cancel()

	client := cryptocompare.NewClient(http.DefaultClient)
	prices, listResponse, err := client.Histoday.Get(ctx, &cryptocompare.HistodayRequest{Fsym: crypto, Tsym: fiat})

	if err != nil || listResponse.StatusCode != 200 {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	response := schema.NewSuccessResponse(prices, "Success", http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}
