package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammed-maher/auth-service/handlers/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("login", Login).Use(middleware.RateLimited())
		v1.POST("register", RegisterUser).Use(middleware.RateLimited())
		protected := v1.Group("protected").Use(middleware.Auth(), middleware.RateLimited())
		{
			protected.GET("ping", Ping)
		}
	}
	return router
}
