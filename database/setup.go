package config

import (
	config "ExchangeRates/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvMongoURI()))
	checkIfErrorLog(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	checkIfErrorLog(err)

	err = client.Ping(ctx, nil)
	checkIfErrorLog(err)

	fmt.Println("Connected to MongoDB Instance.")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("exchange").Collection(collectionName)
	return collection
}

func checkIfErrorLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
