package db

import (
	"log"

	"github.com/iliadril/go-docker-todo/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	dbUrl := "postgres://pg:root@localhost:5432/todos"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
