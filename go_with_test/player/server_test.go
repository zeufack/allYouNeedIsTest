package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores  map[string]int
	winCall []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCall = append(s.winCall, name)
}

func newGetScoreRequest(t *testing.T, name string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostScoreRequest(t *testing.T, name string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func assertStatut(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestGetPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		}, nil,
	}
	server := &PlayerServer{&store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest(t, "Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		assertStatut(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest(t, "Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "10"
		assertStatut(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest(t, "Appollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusNotFound

		assertStatut(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{store: &store}
	// t.Run("it return accept on Post", func(t *testing.T) {
	// 	request := newPostScoreRequest(t, "Pepper")
	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	assertStatut(t, response.Code, http.StatusAccepted)

	// })

	t.Run("it record a win when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostScoreRequest(t, "Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatut(t, response.Code, http.StatusAccepted)
		if len(store.winCall) != 1 {
			t.Errorf("want %d for recording , got %d", 1, len(store.winCall))
		}
		if store.winCall[0] != player {
			t.Errorf("want %q to be store, got %q", player, store.winCall[0])
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	serve := PlayerServer{&store}

	t.Run("it retun 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		serve.ServeHTTP(response, request)

		assertStatut(t, response.Code, http.StatusOK)
	})
}
