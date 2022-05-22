package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/models"
)

func (h handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Read the dynamid id para
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the book by Id
	var todo models.Todo
	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete buk
	h.DB.Delete(&todo)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
