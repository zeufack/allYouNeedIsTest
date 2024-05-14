package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinAndRetreiveThen(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store: store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(t, player))
	assertStatut(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
