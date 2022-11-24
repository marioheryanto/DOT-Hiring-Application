package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgrespw@localhost:34157"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
