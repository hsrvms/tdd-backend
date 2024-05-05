package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store: store}

	log.Println("server is running on 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
