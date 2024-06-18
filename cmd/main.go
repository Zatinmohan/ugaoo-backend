package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zatinmohan007/ugaoo/cmd/database"
	"github.com/zatinmohan007/ugaoo/cmd/router"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading the config from .env file")
		err = godotenv.Load(".env")

		if err != nil {
			log.Printf("Error while loading .env file %v\n", err)
		}
		log.Println(".env loaded successfully")
	}

	_, err := database.InitDatabase()

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := router.SetupRoutes()
	port := os.Getenv("PORT")

	log.Println("Starting server at port ", port)
	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatalln("There's an error with server", err)
	}
}
