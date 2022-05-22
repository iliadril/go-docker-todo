package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/models"
)

// GET method
// fetches Todo object from the db based on given ID
func (h handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	// Read dynamid id para and convert it to an integer value
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find the Todo by Id
	var todo models.Todo
	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// respond to the request
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
