package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/krishnachowdaryvanam/GOLANG/models.git"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func SetUp() {
	host := "localhost"
	port := "5432"
	dbName := "Book"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " dbName=" + dbName + " username=" + username + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
