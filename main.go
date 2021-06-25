package main

import (
	"log"
	"net/http"
	"os"

	h "github.com/jeromewu/alfred-bot/internal/handlers"
	m "github.com/jeromewu/alfred-bot/internal/models"
)

func main() {
	conf := m.NewConf("./config.yaml")

	http.HandleFunc("/webhook", h.Webhook(conf))
	http.HandleFunc("/", h.Root)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Alfred Bot Server Started")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
