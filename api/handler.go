package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	_ "github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"

	"log"
	"net/http"
)

var database = NewService()

//Request ...
type Request struct {
	Url string `json:"name"`
}

type Response struct {
	ID           string `bson:"id"`
	ShortenedURL string `json:"shortenedURL" bson:"url"`
	Hash         string `json:"hash"         bson:"hash"`
	Valid        bool   `json:"isValidURL"`
}

//ShortUrl ...
func ShortUrl(c *fiber.Ctx) error {

	body := new(Request)
	err := c.BodyParser(body)
	if err != nil {
		log.Fatal(err)
	}
	if !govalidator.IsURL(body.Url) {
		res := Response{
			Valid: false,
		}
		return c.JSON(res)
	}

	// c.Response()
	// c.Request()

	database.VerifyHash("xxx")

	c.Format("<h1>Maximoooo</h1>")
	return c.SendStatus(http.StatusCreated)
}
