package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mohammed-maher/auth-service/config"
	"github.com/mohammed-maher/auth-service/database"
	"github.com/mohammed-maher/auth-service/handlers"
	"log"
)

func main() {

	// load config
	conf := config.Load()
	log.Println("establishing db connection at: ", conf.DB_DSN)

	// connect to database
	if err := database.Connect(conf.DB_DSN); err != nil {
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
	err = r.Run(":" + conf.PORT)

	if err != nil {
		log.Fatal("failed to start server on port", conf.PORT, err)
	}

}
