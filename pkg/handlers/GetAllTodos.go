package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iliadril/go-docker-todo/pkg/models"
)

// GET method
// fetches ALL Todo objects from the db
func (h handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	// create slice to hold Todo objects
	var todos []models.Todo
	// add todo objects to the todos slice, throw an error if unsuccessful
	if result := h.DB.Find(&todos); result.Error != nil {
		log.Println(result.Error)
	}

	// respond to the request
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
