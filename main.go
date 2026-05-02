package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()
	handler := NewHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("/shorten",handler.ShortenURL)
	mux.HandleFunc("/",handler.Redirect)

	log.Println("Server Running of http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090",mux))
}