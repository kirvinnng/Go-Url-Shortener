package api

import (
	"context"

	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Service ...
type MongoHandle struct {
	*mongo.Collection
}

type MongoGateway interface {
	InsertUrl()
	UpdateUrl()
	VerifyHash(hashToFind string) bool
}

type Service struct {
	MongoGateway
}

func NewService() MongoGateway {

	clie := db.InitMongoDB()
	mg := &MongoHandle{clie}
	return Service{mg}
}

func (h *MongoHandle) InsertUrl() {

	//name to the database

	//collName := h.Name()

	//collection := h.Database().Collection(collName)

}

func (h *MongoHandle) UpdateUrl() {

}

func (h *MongoHandle) VerifyHash(hashToFind string) bool {

	ctx := context.Background()
	var result bson.D
	err := h.FindOne(ctx, bson.D{{"hash", hashToFind}}).Decode(&result)

	if err != nil {
		return false
	}
	return true
}
