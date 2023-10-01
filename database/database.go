package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(url string) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverApi)
	database, err := mongo.Connect(context.TODO(), clientOptions)
	client = database
	return err
}

func GetClient() *mongo.Client {
	return client
}
