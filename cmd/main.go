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
	r.HandleFunc("/todos/{id}", h.GetTodo).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.AddTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", h.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", h.DeleteTodo).Methods(http.MethodDelete)

	log.Println("Starting server :8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
