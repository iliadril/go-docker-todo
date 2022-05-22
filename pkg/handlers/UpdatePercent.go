package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/iliadril/go-docker-todo/pkg/models"
)

// PUT method
// Updates a Todo object completeness % based on given id
// completeness value gets updated with the PUT {perc: [value]}, other columns remain the same
// if ?perc param is equal to done (case insensitive) PUT data is overriten to 100 which marks Todo as Done
func (h handler) UpdatePercent(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// get the completeness % from query param
	perc := r.FormValue("perc")

	// Read request body and make sure to close the request when function returns
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// create temporary todo object from unmarshalled request
	var updatedTodo models.Todo
	json.Unmarshal(body, &updatedTodo)

	// update the Todo based on given Id
	if strings.EqualFold("done", perc) {
		updatedTodo.Completeness = 100
	}
	h.DB.Model(&models.Todo{Id: id}).Update("Completeness", updatedTodo.Completeness)

	// respond to the request
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}
