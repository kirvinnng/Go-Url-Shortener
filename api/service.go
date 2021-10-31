package api

import (
	"context"
	"fmt"
	"log"

	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Service ...
type MongoHandle struct {
	*mongo.Collection
}

type MongoGateway interface {
	InsertUrl(document ResponseClient)

	GetDocument(url string) (string, string)

	VerifyUrl(url string) bool

	VerifyHash(hashToFind string) bool

	SearchHash(hash string) (string, bool)
}

type Service struct {
	MongoGateway
}

func NewService() Service {

	clie := db.InitMongoDB()
	mg := &MongoHandle{clie}
	return Service{mg}
}

func (h *MongoHandle) InsertUrl(document ResponseClient) {

	ctx := context.Background()
	data := struct {
		Hash         string `bson:"hash"`
		OriginalUrl  string `bson:"url"`
		ShortenedUrl string `bson:"url_hash"`
	}{
		Hash:         document.Hash,
		OriginalUrl:  document.OriginalUrl,
		ShortenedUrl: document.ShortenedURL,
	}

	coll, err := h.InsertOne(ctx, data)

	if err != nil {
		log.Println(err)

	}

	log.Println("Inserted a single document: ", coll.InsertedID)

}

//GetDocument returns the hash & the shortedUrl corresponding to the original url
func (h *MongoHandle) GetDocument(url string) (string, string) {

	ctx := context.Background()
	var result bson.D
	err := h.FindOne(ctx, bson.D{primitive.E{Key: "url", Value: url}}).Decode(&result)
	if err != nil {
		log.Fatal(err)

	}

	mp := result.Map()

	justHash := fmt.Sprint(mp["hash"])
	urlWithHash := fmt.Sprint(mp["url_hash"])

	return justHash, urlWithHash
}

func (h *MongoHandle) SearchHash(hash string) (string, bool) {

	ctx := context.Background()
	var result bson.D
	err := h.FindOne(ctx, bson.D{primitive.E{Key: "hash", Value: hash}}).Decode(&result)
	if err != nil {
		return "", false
	}

	mp := result.Map()
	originalUrl := fmt.Sprint(mp["url"]) // this is nil!!

	fmt.Println(originalUrl)

	return originalUrl, true
}

func (h *MongoHandle) VerifyUrl(url string) bool {

	ctx := context.Background()
	var result bson.D
	err := h.FindOne(ctx, bson.D{primitive.E{Key: "url", Value: url}}).Decode(&result)

	return err == nil
}

func (h *MongoHandle) VerifyHash(hashToFind string) bool {

	ctx := context.Background()
	var result bson.D
	err := h.FindOne(ctx, bson.D{primitive.E{Key: "hash", Value: hashToFind}}).Decode(&result)

	return err == nil
}
