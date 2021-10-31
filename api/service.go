package api

import (
	"context"

	_ "github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Service ...
type MongoHandle struct {
	Mongo *mongo.Collection
}

type Service struct {
	user Gateway
}

type Gateway interface {
	InsertUrl()
	UpdateUrl()
	VerifyHash(hashToFind string) bool
}

// func NewService() Service {

// 	cl := db.InitMongoDB()

// 	return Service{user: cl}
// }

func (h *MongoHandle) InsertUrl() {

	//name to the database

	//collName := h.Name()

	//collection := h.Database().Collection(collName)

}

func UpdateUrl() {

}

func (h *MongoHandle) VerifyHash(hashToFind string) bool {

	ctx := context.Background()
	var result bson.D
	err := h.Mongo.FindOne(ctx, bson.D{{"hash", hashToFind}}).Decode(&result)

	if err != nil {
		return false
	}
	return true
}
