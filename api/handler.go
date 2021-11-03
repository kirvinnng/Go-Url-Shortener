package api

import (
	"log"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/hash"
)

var database = NewService()

//RequestClient ...
type RequestClient struct {
	Url string `json:"url"`
}

type ResponseClient struct {
	Hash         string `json:"hash"          bson:"hash"    `
	OriginalUrl  string `json:"originalURL"  bson:"url"     `
	ShortenedURL string `json:"shortenedURL"  bson:"url_hash"`
	Valid        bool   `json:"isValidURL"                   `
}

func ShortUrl(c *fiber.Ctx) error {

	body := new(RequestClient)
	err := c.BodyParser(body) //* get the request url
	if err != nil {
		log.Fatal(err)
	}

	if !govalidator.IsURL(body.Url) { //* check if the url is valid
		res := ResponseClient{
			Valid: false,
		}
		err := c.SendStatus(http.StatusNotAcceptable)
		if err != nil {
			log.Print(err)
		}
		return c.JSON(res)
	}
	// If the url doesn't start with http: its added
	parseURL, _ := url.Parse(body.Url)
	if parseURL.Scheme == "" {
		body.Url = "http://" + body.Url
	}

	//* check if the url is valid in the database
	if database.VerifyUrl(body.Url) {

		//* if the url exists, the existing data is returned
		res := sendExistingUrlWithHash(c, body.Url)

		return c.JSON(res)

	} else {

		//* if the url not exists
		res := createNewUrlWithHash(c, body.Url)
		err := c.SendStatus(http.StatusCreated)
		if err != nil {
			log.Print(err)
		}
		return c.JSON(res)
	}
}

func sendExistingUrlWithHash(c *fiber.Ctx, url string) ResponseClient {

	//* get the hash & the url with the hash corresponding to the url
	justHash, urlWithHash := database.GetDocument(url)

	res := ResponseClient{
		Valid:        true,
		Hash:         justHash,
		ShortenedURL: urlWithHash,
		OriginalUrl:  url,
	}
	//* if the url exists, the existing data is returned
	return res

}

func createNewUrlWithHash(c *fiber.Ctx, originalUrl string) ResponseClient {

	uriHash := hash.GenerateHash(6)

	//* check if the generated hash isn't in the database
	for database.VerifyHash(uriHash) {
		uriHash = hash.GenerateHash(6)

	}

	//* Insert this new url with the hash

	resp := ResponseClient{
		Hash:         uriHash,
		OriginalUrl:  originalUrl,
		ShortenedURL: c.BaseURL() + "/" + uriHash,
		Valid:        true,
	}

	database.InsertUrl(resp)
	return resp
}

func Redirect(c *fiber.Ctx) error {

	uriHash := c.Params("hash")

	originalUrl, exist := database.SearchHash(uriHash)

	if exist {
		return c.Redirect(originalUrl, http.StatusMovedPermanently)
	}

	return c.SendStatus(http.StatusNotFound)
}

func Index(c *fiber.Ctx) error {
	return c.SendFile("./template/index.html")
}
