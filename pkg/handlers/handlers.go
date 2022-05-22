package handlers

import "gorm.io/gorm"

// create new handler struct to hold the DB
type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
