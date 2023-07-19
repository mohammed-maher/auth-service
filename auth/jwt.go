package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (tokenString string, err error) {
	exp := time.Now().Add(30 * time.Second)
	claims := JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)

	return
}

func ValidateToken(tokenString string) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	claims := token.Claims.(*JWTClaim)
	if claims.ExpiresAt < time.Now().Unix() {
		return errors.New("token expired")
	}
	return
}
