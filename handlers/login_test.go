package handlers

import (
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/mohammed-maher/auth-service/config"
	"github.com/mohammed-maher/auth-service/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	godotenv.Load(".env.test")
}

func TestLoginWithoutCredentials(t *testing.T) {
	t.Log("given a user is trying to login without supplying credentials, user should receive bad request")
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginWithInvalidEmail(t *testing.T) {
	t.Log("given a user is trying to login supplying non-existing email, user should receive not found ")
	router := SetupRouter()

	w := httptest.NewRecorder()
	conf := config.Load()
	err := database.Connect(conf.DB_DSN)
	if err != nil {
		t.Log("could not connect to db", err)
		t.Fail()
	}
	bodyReader := strings.NewReader(`{"email": "doesnotexist@email.com", "password": ""}`)
	req, _ := http.NewRequest("POST", "/api/v1/login", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestLoginWithValidCredentials(t *testing.T) {
	t.Log("given a user is trying to login supplying valid credentials, user should ok status and a token")
	router := SetupRouter()

	w := httptest.NewRecorder()
	conf := config.Load()
	err := database.Connect(conf.DB_DSN)
	if err != nil {
		t.Log("could not connect to db", err)
		t.Fail()
	}

	user := database.User{
		Name:     "test user",
		Username: "test user",
		Email:    "test@user.com",
		Password: "testpass",
	}

	user.HashPassword(user.Password)

	bodyReader := strings.NewReader(`{"email": "test@user.com", "password": "testpass"}`)
	req, _ := http.NewRequest("POST", "/api/v1/login", bodyReader)

	database.CreateUser(user)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if !strings.Contains(w.Body.String(), "token") {
		t.Log("user login response does not contain token")
		t.Fail()
	}

	// @TODO clean up
}
