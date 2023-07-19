package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mohammed-maher/auth-service/database"
	"log"
	"os"
)

func main() {
	connectionString := os.Getenv("DB_DSN")
	log.Println("establishing db connection at: ", connectionString)

	if err := database.Connect(os.Getenv("DB_DSN")); err != nil {
		log.Fatal("failed to establish db connection: ", err)
	}

	log.Println("migrating db..")
	if err := database.Migrate(); err != nil {
		log.Println("failed to migrate db with error ", err)
	}

}
