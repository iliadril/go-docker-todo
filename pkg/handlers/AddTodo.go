package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iliadril/go-docker-todo/pkg/models"
)

// POST method
// Adds Todo object to the DB based on POST data
func (h handler) AddTodo(w http.ResponseWriter, r *http.Request) {
	// make sure to close the body when function returns
	defer r.Body.Close()

	// read body, throw error if unsuccessfull
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// create new Todo object from unmarshalled request data
	var todo models.Todo
	json.Unmarshal(body, &todo)

	// add todo object to the DB, throw an error if unsuccessfull
	if result := h.DB.Create(&todo); result.Error != nil {
		log.Println(result.Error)
	}

	// respond to the request
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
