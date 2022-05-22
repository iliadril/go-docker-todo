package models

import "time"

// create Todo struct to serve as a base model for the app
type Todo struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Expiry       time.Time `json:"exp"`
	Title        string    `json:"title"`
	Describtion  string    `json:"desc"`
	Completeness int       `json:"perc"`
}
