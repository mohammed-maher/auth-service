package handlers

import (
	"github.com/go-playground/assert/v2"
	"github.com/mohammed-maher/auth-service/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimitedWithAuthorization(t *testing.T) {
	t.Log("given a user hits requests limit, user should receive too many requests status")
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/api/v1/protected/ping", nil)
	token, _ := auth.GenerateJWT("test@user.com")
	req.Header.Add("Authorization", token)

	for i := 0; i <= 10; i++ {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
	}
	check := httptest.NewRecorder()
	router.ServeHTTP(check, req)
	assert.Equal(t, http.StatusTooManyRequests, check.Code)
}

func TestRateLimitedWithoutAuthorization(t *testing.T) {
	t.Log("given a user hits requests limit, user should receive too many requests status")
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/api/v1/protected/ping", nil)

	for i := 0; i <= 10; i++ {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
	}
	check := httptest.NewRecorder()
	router.ServeHTTP(check, req)
	assert.Equal(t, http.StatusTooManyRequests, check.Code)
}
