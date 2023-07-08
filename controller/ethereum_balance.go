package controller

import (
	errorHandler "ExchangeRates/error_handling"
	schema "ExchangeRates/schema/base"
	"context"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func GetBalanceUsingAddress(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	address := c.Params("address")
	defer cancel()

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/")

	if err != nil {
		return errorHandler.HandleError(c, http.StatusConflict, err)
	}

	addressBytes := convertStringToByteArray(address)
	balance, err := client.BalanceAt(ctx, common.Address{addressBytes}, nil)
	if err != nil {
		return errorHandler.HandleError(c, http.StatusBadRequest, err)
	}

	response := schema.NewSuccessResponse(balance, "Success", http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}

func convertStringToByteArray(input string) byte {
	byte_array := []byte(input)[0]
	return byte_array
}
