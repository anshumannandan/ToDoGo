package initializers

import (
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(){

	var err error

	db_path := os.Getenv("DB_PATH")
	db, err = gorm.Open(sqlite.Open(db_path), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

}