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

func SaveNumber(number int32) {
	collection, error := connect_to_table()
	if error != nil {
		panic(error)
	}

	doc := NumberStruct{Number: number}

	_, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
}

func connect_to_table() (*mongo.Collection, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://admin:admin123@taller-admins.ez0xrnf.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client2, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	collection := client2.Database("numbers").Collection("numbers")
	return collection, nil
}

func GetNumbers() []int32 {
	var result []int32
	collection, error := connect_to_table()
	if error != nil {
		panic(error)
	}

	numbers, error := collection.Distinct(context.Background(), "number", bson.D{})
	if error != nil {
		panic(error)
	}
	for _, number := range numbers {
		result = append(result, number.(int32))
	}
	return result
}
