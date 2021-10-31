package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/maximo-torterolo-ambrosini/Go-Url-Shortener/db"
	//"go.mongodb.org/mongo-driver" // DRIVER to install
)

func main() {
	// Setting enviorment variables🔥
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitMongoDB()
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

	app.Get("/", func(c *fiber.Ctx) error {
		c.Format("<h1>Maximoooo</h1>")
		return c.SendStatus(http.StatusCreated)
	})

	app.Listen(PORT)
}
