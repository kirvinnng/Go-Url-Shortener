package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/api"
	//"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"
	// "github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/hash"
)

func main() {
	// Setting enviorment variables🔥
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	// Configuring port
	PORT := ":" + os.Getenv("PORT")
	if PORT == ":" {
		log.Println("PORT env variable is null setting port to 8080")
		PORT += "8080"
	}

	// Initializing Fiber App
	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiber",
		AppName:      "Url Shortener",
	})

	app.Use(logger.New())

	//app.Static()

	// api.SetRoutes(app)

	app.Listen(PORT)
}
