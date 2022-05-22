package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iliadril/go-docker-todo/pkg/models"
)

func (h handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	if result := h.DB.Find(&todos); result.Error != nil {
		log.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
