package mapper

import (
	"ExchangeRates/model"
	response "ExchangeRates/schema/exchange_rates"

	cryptocompare "github.com/lucazulian/cryptocomparego"
	"go.mongodb.org/mongo-driver/mongo"
)

func ToDto(price cryptocompare.Price) response.ExchangePrice {
	return response.ExchangePrice{
		Name:  price.Name,
		Value: price.Value,
	}
}

func ToDtoList(prices []cryptocompare.Price, result mongo.InsertManyResult) response.ExchangeResponse {
	var exchangePriceList []response.ExchangePrice
	var exchangeResponse response.ExchangeResponse
	for i := 0; i < len(prices); i++ {
		var exchangePriceItem response.ExchangePrice
		exchangePriceItem.Name = prices[i].Name
		exchangePriceItem.Value = prices[i].Value
		exchangePriceList = append(exchangePriceList, exchangePriceItem)
	}
	exchangeResponse.ExchangePrice = exchangePriceList
	exchangeResponse.InsertedIDs = result.InsertedIDs
	return exchangeResponse
}

func ToEntityList(prices []cryptocompare.Price) []model.ExchangeRate {
	var exchangePriceList []model.ExchangeRate
	for i := 0; i < len(prices); i++ {
		var exchangePriceItem model.ExchangeRate
		exchangePriceItem.Name = prices[i].Name
		exchangePriceItem.Value = prices[i].Value
		exchangePriceList = append(exchangePriceList, exchangePriceItem)
	}
	return exchangePriceList
}
