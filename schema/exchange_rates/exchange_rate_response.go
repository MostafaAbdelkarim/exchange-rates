package model

type ExchangePrice struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type ExchangeResponse struct {
	ExchangePrice []ExchangePrice `json:"exchangePrices"`
	InsertedIDs   []interface{}   `json:"insertedIds"`
}
