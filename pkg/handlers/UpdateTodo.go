package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/models"
)

func (h handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedTodo models.Todo
	json.Unmarshal(body, &updatedTodo)

	// update
	var todo models.Todo

	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	todo.Expiry = updatedTodo.Expiry
	todo.Title = updatedTodo.Title
	todo.Describtion = updatedTodo.Describtion
	todo.Completeness = updatedTodo.Completeness

	h.DB.Save(&todo)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}
