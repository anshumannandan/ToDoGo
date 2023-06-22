package initializers

import (
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){

	var err error

	db_path := os.Getenv("DB_PATH")
	DB, err = gorm.Open(sqlite.Open(db_path), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

}