package handlers

import (
	"github.com/go-playground/assert/v2"
	"github.com/mohammed-maher/auth-service/auth"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingWithoutAuth(t *testing.T) {
	t.Log("given an unauthorized user is trying to call protected route, user should receive 401")
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/protected/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestPingWithAuth(t *testing.T) {
	t.Log("given an authorized user is trying to call protected route, user should receive status 200 and pong")
	router := SetupRouter()
	w := httptest.NewRecorder()
	token, _ := auth.GenerateJWT("testuser@mail.co")
	req, _ := http.NewRequest("GET", "/api/v1/protected/ping", nil)
	req.Header.Add("Authorization", token)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	if !strings.Contains(w.Body.String(), "pong") {
		t.Log("user should receive pong response")
		t.Fail()
	}
}
