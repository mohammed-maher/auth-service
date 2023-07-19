package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mohammed-maher/auth-service/database"
	"github.com/mohammed-maher/auth-service/handlers"
	"log"
	"os"
)

func main() {

	connectionString := os.Getenv("DB_DSN")
	log.Println("establishing db connection at: ", connectionString)

	// connect to database
	if err := database.Connect(os.Getenv("DB_DSN")); err != nil {
		log.Fatal("failed to establish db connection: ", err)
	}

	// migrate db
	log.Println("migrating db..")
	if err := database.Migrate(); err != nil {
		log.Println("failed to migrate db with error ", err)
	}

	// listen and serve
	r := handlers.SetupRouter()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Println("failed to set router trusted proxies")
	}
	err = r.Run(":9000")
	if err != nil {
		log.Fatal("failed to start server ", err)
	}

}
