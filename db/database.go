package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Shortened struct {
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

//MongoCollection ...
type MongoCollection struct {
	CollectionDB *mongo.Collection
}

// InitMongoDB ...
func InitMongoDB() *MongoCollection {

	env := Root()

	clientOptions := options.Client().ApplyURI(env.uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	pingDB(client)
	cl := client.Database(env.name).Collection("information")

	return &MongoCollection{cl}
}

func pingDB(db *mongo.Client) {
	err := db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
