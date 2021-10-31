package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//ConfigRoot ...
type ConfigRoot struct {
	Name string
	Uri  string
}

//Root ...
func Root() ConfigRoot {

	// Setting enviorment variables🔥
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	URI := os.Getenv("DB_URI")
	DB := os.Getenv("DB_NAME")

	return ConfigRoot{
		Name: DB,
		Uri:  URI,
	}
}
