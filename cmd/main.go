package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/db"
	"github.com/iliadril/go-docker-todo/pkg/handlers"
)

func main() {
	DB := db.Initialize()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/todos", h.GetAllTodos).Methods(http.MethodGet)

	log.Println("Starting server :8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
