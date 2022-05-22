package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/models"
)

// DELETE method
// Deletes a Todo object from the DB based on its id
func (h handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Read the dynamid id para and convert it to an integer value
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the Todo by Id
	var todo models.Todo
	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete the Todo
	h.DB.Delete(&todo)

	// respond to the request
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Deleted")
}
