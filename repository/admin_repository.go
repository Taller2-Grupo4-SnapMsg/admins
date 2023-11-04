package repository

import (
	"admins/structs"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NumberStruct struct {
	Number int32
}

var (
	collection_global *mongo.Collection
	connected         bool
)

func Connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	collection_global = client.Database("admins").Collection("admins")
	connected = true
}

func ConnectToNumbers() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	collection_global = client.Database("numbers").Collection("numbers")
	connected = true
}

/**
 * This function saves an admin on the database.
 * @param admin The admin to be saved.
 * @return the admin if it was saved, nil otherwise.
 */
func SaveAdmin(admin *structs.AdminStruct) (*structs.AdminStruct, error) {
	if !connected {
		Connect()
	}
	_, err := collection_global.InsertOne(context.Background(), admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

/**
 * This function gets an admin from the database.
 * @param email The email of the admin.
 * @return the admin if it was found, nil otherwise.
 */
func GetAdmin(email string) *structs.AdminStruct {
	if !connected {
		Connect()
	}
	var admin structs.AdminStruct
	filter := bson.D{{Key: "email", Value: email}}
	err := collection_global.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// manage not found
			return nil
		}
		// real error
		panic(err)
	}
	return &admin
}

/**
 * This function deletes an admin from the database.
 * @param email The email of the admin.
 * @return nil if everything was ok, error otherwise.
 */
func DeleteAdmin(email string) (int64, error) {
	if !connected {
		Connect()
	}
	filter := bson.D{{Key: "email", Value: email}}
	delete_result, err := collection_global.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	return delete_result.DeletedCount, nil
}

func SaveNumber(number int32) {
	if !connected {
		ConnectToNumbers()
	}
	doc := NumberStruct{Number: number}

	_, err := collection_global.InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
}

func GetNumbers() []int32 {
	if !connected {
		ConnectToNumbers()
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
