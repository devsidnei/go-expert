package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Server: handler started")

	defer log.Println("Server: handler ended")

	select {
	case <-time.After(5 * time.Second):
		log.Printf("Server: request processed")
		w.Write([]byte("Request processed"))
	case <-ctx.Done():
		log.Printf("Server: request cancelled")
	}
}
