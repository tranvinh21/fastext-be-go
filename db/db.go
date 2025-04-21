package db

import (
	"log"

	"github.com/tranvinh21/fastext-be-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.Envs.DB.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return db
}
