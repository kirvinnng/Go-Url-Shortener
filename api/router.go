package api

import (
	"github.com/gofiber/fiber/v2"
)

//SetRoutes ...
func SetRoutes(app *fiber.App) {

	//app.Get("/hash")
	//app.Get("/")
	app.Post("/", ShortUrl)

}
