package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExchangeRate struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Value float64            `json:"value,omitempty" validate:"required"`
}

type ExchangeObject struct {
	Cryptocurrency string         `json:"cryptoCurrecny"`
	ExchangeRate   []ExchangeRate `json:"exchangeRate"`
}
