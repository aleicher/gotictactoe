package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}

func TestHandleGamesIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/games", nil)
	response := httptest.NewRecorder()

	GamesHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}
