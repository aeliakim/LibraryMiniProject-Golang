package main

import (
	"LibraryMiniProject-Golang/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ServeHome).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
