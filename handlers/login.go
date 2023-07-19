package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammed-maher/auth-service/auth"
	"github.com/mohammed-maher/auth-service/database"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginRequest LoginRequest
	var user database.User

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incompatible request"})
		ctx.Abort()
		return
	}

	// verify supplied credentials
	record := database.Instance.Where("email = ?", loginRequest.Email).First(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(loginRequest.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	// generate and return token
	tokenString, err := auth.GenerateJWT(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
