package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func InitDatabase() (*mongo.Client, error) {

	if client != nil {
		fmt.Println("Returning mongoDb Client")
		return client, nil
	}

	log.Println("Connecting to MongoDB....")
	uri := os.Getenv("MONGO_HOST")

	option := options.Client().ApplyURI(uri)
	ctx := context.TODO()

	client, err = mongo.Connect(ctx, option)

	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected Successfully")
	return client, nil
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	if client == nil {
		return nil, errors.New("please connect to the client and then try again")
	}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName)

	return collection, nil
}
