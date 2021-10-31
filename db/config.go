package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ConfigRoot struct {
	name string
	uri  string
}

func Root() ConfigRoot {

	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	URI := os.Getenv("DB_URI")
	DB := os.Getenv("DB_NAME")

	return ConfigRoot{
		name: DB,
		uri:  URI,
	}
}
