package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	//server := &PlayerServer{NewInMemoryPlayerStore()}

	//server.store.RecordWin("Pepper")

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could no listen on port 5000 %v", err)
	}
}
