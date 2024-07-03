package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() (*mongo.Database, error) {

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		// AuthSource:    "WPAPDB",
		Username: "dexter",
		Password: "Cetnbcm88Cetnbcm88$",
		// Password: "Cetnbcm88Cetnbcm88$$Cetnbcm88Cetnbcm88$$",
	}

	// credential := options.Credential{
	// 	AuthMechanism: "SCRAM-SHA-256",
	// 	AuthSource:    "WPAPDB",
	// 	Username:      "Dexter",
	// 	Password:      "Cetnbcm88Cetnbcm88",
	// }

	// Set client options
	clientOptions := options.Client().ApplyURI(MongoUrl).SetAuth(credential)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client.Database(MongoDatabase), nil
}
