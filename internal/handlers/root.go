package handlers

import (
	"fmt"
	"net/http"
	"time"
)

var start = time.Now()

func get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Started at %v, running for %v already", start, time.Now().Sub(start))))
}

func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
	} else {
		http.Error(w, "405 METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}
}
