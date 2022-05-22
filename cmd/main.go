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
	// Initialize the DataBase connection
	DB := db.Initialize()
	h := handlers.New(DB)
	// Initialize the mux Router
	r := mux.NewRouter()

	// Initialize Routes
	// GET methods
	r.Path("/todos").Queries("range", "{range}").HandlerFunc(h.GetByDate).Methods(http.MethodGet) // fetches incoming Todos
	r.HandleFunc("/todos", h.GetAllTodos).Methods(http.MethodGet)                                 // fetches ALL Todos
	r.HandleFunc("/todos/{id:[0-9]+}", h.GetTodo).Methods(http.MethodGet)                         // fetches Todos based on given Id
	// POST methods
	r.HandleFunc("/todos", h.AddTodo).Methods(http.MethodPost) // add a new Todo to the db
	// PUT methods
	r.Path("/todos/{id:[0-9]+}").Queries("perc", "{perc}").HandlerFunc(h.UpdatePercent).Methods(http.MethodPut) // updates Todos completeness % based on corresponding Id
	r.HandleFunc("/todos/{id:[0-9]+}", h.UpdateTodo).Methods(http.MethodPut)                                    // updates EVERY value of a Todo based on given Id
	// DELETE methods
	r.HandleFunc("/todos/{id:[0-9]+}", h.DeleteTodo).Methods(http.MethodDelete) // deletes a Todo object based on given Id

	// Initializing the server on port 8080
	log.Println("Starting server :8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second, // Set Timeouts against SlowIoris
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
