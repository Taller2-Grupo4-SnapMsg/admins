package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NumberStruct struct {
	Number int32
}

var (
	client_global     *mongo.Client
	collection_global *mongo.Collection
	connected         bool
)

func Connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://admin:admin123@taller-admins.ez0xrnf.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	collection_global = client.Database("numbers").Collection("numbers")
	connected = true
}

func SaveNumber(number int32) {
	if !connected {
		Connect()
	}
	doc := NumberStruct{Number: number}

	_, err := collection_global.InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
}

func GetNumbers() []int32 {
	if !connected {
		Connect()
	}
	var result []int32
	numbers, error := collection_global.Distinct(context.Background(), "number", bson.D{})
	if error != nil {
		panic(error)
	}
	for _, number := range numbers {
		result = append(result, number.(int32))
	}
	return result
}
