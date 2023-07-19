package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mohammed-maher/auth-service/database"
	"github.com/mohammed-maher/auth-service/handlers"
	"github.com/mohammed-maher/auth-service/handlers/middleware"
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

	// run webserver
	if err := Serve(); err != nil {
		log.Fatal("failed to init server: ", err)
	}

}

func Serve() error {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("login", handlers.Login)
		v1.POST("register", handlers.RegisterUser)
		protected := v1.Group("protected").Use(middleware.Auth())
		{
			protected.GET("ping", handlers.Ping)
		}
	}

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return err
	}

	return router.Run(":9000")
}
