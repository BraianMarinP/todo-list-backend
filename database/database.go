package database

import (
	"log"
	"todo-list-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dataSourceName := config.GetEnvVariable("DB_DSN")
	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database. ", err)
	}
}
