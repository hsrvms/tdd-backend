package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store: store}

	log.Println("server is running on 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
