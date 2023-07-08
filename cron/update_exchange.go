package cron

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/robfig/cron"

	cryptocompare "github.com/lucazulian/cryptocomparego"
)

func UpdateExchangeRatesEvery5Minutes() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	client := cryptocompare.NewClient(http.DefaultClient)
	cronJob := cron.New()
	defer cancel()

	cronJob.AddFunc("*/5 * * * *", func() {
		var all []string
		var pricesArray []cryptocompare.Price
		all = append(all, "USD")
		prices, _, _ := client.Price.List(ctx, &cryptocompare.PriceRequest{Fsym: "BTC", Tsyms: all})
		pricesArray = append(pricesArray, prices...)
		fmt.Println("Cron Job is running -- ")
	})
	cronJob.Start()

}
