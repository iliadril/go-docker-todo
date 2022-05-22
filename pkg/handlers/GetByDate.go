package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/iliadril/go-docker-todo/pkg/models"
)

// GET method
// fetches Todo object from the db based on the query ?range=[today,tomorrow,week] param
func (h handler) GetByDate(w http.ResponseWriter, r *http.Request) {
	// get the range from param
	date := r.FormValue("range")

	// create slice to hold Todo objects
	var todos []models.Todo
	t := time.Now() // get current time

	// add corresponding Todo objects to the slice based on the ?range param
	switch {
	case strings.EqualFold("today", date): // today's date tasks
		if result := h.DB.Where("DATE(expiry) = ?", t.Format("2006-01-02")).Find(&todos); result.Error != nil {
			log.Println(result.Error)
		}
	case strings.EqualFold("tomorrow", date): // tomorrow's date tasks
		if result := h.DB.Where("DATE(expiry) = ?", t.Add(24*time.Hour).Format("2006-01-02")).Find(&todos); result.Error != nil {
			log.Println(result.Error)
		}
	case strings.EqualFold("week", date): // all tasks for the week
		if result := h.DB.Where("expiry BETWEEN ? AND ?", t.Format("2006-01-02"), t.Add(24*7*time.Hour).Format("2006-01-02")).Find(&todos); result.Error != nil {
			log.Println(result.Error)
		}
	default: // add all todos if ?range doesn't match
		if result := h.DB.Find(&todos); result.Error != nil {
			log.Println(result.Error)
		}
	}

	// respond to the request
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
