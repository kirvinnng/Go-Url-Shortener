package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {

	app.Get("/:hash", GetShortUrl)
	app.Post("/")

}
