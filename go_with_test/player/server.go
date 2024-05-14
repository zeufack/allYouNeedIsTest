package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	player := strings.TrimPrefix(request.URL.Path, "/players/")
	switch request.Method {
	case http.MethodPost:
		p.processWins(response, player)

	case http.MethodGet:
		p.showScore(response, player)
	}

}

func (p *PlayerServer) showScore(response http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		response.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(response, score)
}

func (p *PlayerServer) processWins(response http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	response.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(player string) string {
	if player == "peppers" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}
	return " "
}
