package main

import (
	"net/http"
	
	router "github.com/FauzanAr/clean-and-go/router"
)

func main() {
	PORT := ":9000"

	// Router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := []byte("Server up and running")
		w.Write(message)
	})
	
	routes := http.HandlerFunc(router.Serve)
	http.ListenAndServe(PORT, routes)
}