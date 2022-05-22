package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iliadril/go-docker-todo/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Todo
	json.Unmarshall(body, &todo)

	if result := h.DB.Create(&todo); result.Error != nil {
		log.Println(result.Error)
	}

	w.WtireHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
