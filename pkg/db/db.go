package db

import (
	"log"

	"github.com/iliadril/go-docker-todo/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize() *gorm.DB {
	// set connection URL
	dbUrl := "postgres://pg:root@db:5432/todos"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logger to display Raw SQL Queries
	})
	if err != nil {
		log.Fatalln(err)
	}

	// automatically migrate the schema (create tables, missing foreign keys, etc.) based on Todo struct
	// used to create the table as well as keep it up to date (unused columns WON'T be deleted!)
	db.AutoMigrate(&models.Todo{})

	return db
}
