package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB ...
func InitMongoDB() *mongo.Collection {

	env := Root()

	clientOptions := options.Client().ApplyURI(env.Uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	pingDB(client)
	cl := client.Database(env.Name).Collection("information")
	return cl
}

func pingDB(db *mongo.Client) {
	err := db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
