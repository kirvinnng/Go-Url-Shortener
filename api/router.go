package api

import (
	"github.com/gofiber/fiber/v2"
)

//SetRoutes ...
func SetRoutes(app *fiber.App) {
	app.Get("/", Index)
	app.Get("/:hash", Redirect)
	app.Post("/short", ShortUrl)

}
