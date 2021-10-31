package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/hash"

	"github.com/gofiber/fiber/v2"

	"log"
	"net/http"
)

var database = NewService()

//Request ...
type Request struct {
	Url string `json:"url"`
}

type Response struct {
	ShortenedURL string `json:"shortenedURL" `
	Hash         string `json:"hash"         `
	Valid        bool   `json:"isValidURL"`
}

//ShortUrl ...
func ShortUrl(c *fiber.Ctx) error {

	body := new(Request)
	err := c.BodyParser(body) //* get the request url
	if err != nil {
		log.Fatal(err)
	}

	if !govalidator.IsURL(body.Url) { //* check if the url is valid
		res := Response{
			Valid: false,
		}
		return c.JSON(res)
	}

	//* check if the url is valid in the database
	if database.VerifyUrl(body.Url) {

		//* if the url exists, the existing data is returned
		res := sendExistingUrlWithHash(c, body.Url)

		return c.JSON(res)

	} else {

		//* if the url not exists
		createNewUrlWithHash(body.Url)

		return c.SendStatus(http.StatusCreated)

	}
}

func sendExistingUrlWithHash(c *fiber.Ctx, url string) Response {

	//* get the hash & the url with the hash corresponding to the url
	justHash, urlWithHash := database.GetDocument(url)

	res := Response{
		Valid:        true,
		Hash:         justHash,
		ShortenedURL: urlWithHash,
	}
	//* if the url exists, the existing data is returned
	return res

}

func createNewUrlWithHash(originalUrl string) {

	uriHash := hash.GenerateHash(6)

	//* check if the generated hash isn't in the database
	for database.VerifyHash(uriHash) {
		uriHash = hash.GenerateHash(6)

	}

	//* Insert this new url with the hash
	resp := ResponseMongo{
		Hash:         uriHash,
		OriginalUrl:  originalUrl,
		ShortenedURL: "maxi.com" + "/" + uriHash,
	}

	database.InsertUrl(resp)
}
