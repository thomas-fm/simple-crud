package config

import (
	"fmt"
	"log"
	"os"
	"test-yukbisnis/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const (
// 	dbUser = "root"
// 	dbPass = ""
// 	dbHost = "localhost"
// 	dbName = "dorayaki_api"
// )

var (
	dbUser string
	dbPass string
	dbHost string
	dbName string
	dbPort string
)

func SetupDB() (db *gorm.DB) {
	var err error
	err = godotenv.Load()

	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")

	// dbHost = "'%'"

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database oo")
	}

	// migrate models to database
	db.Debug().AutoMigrate(&models.CV{}, &models.Contact{}, &models.Education{}, &models.Experience{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
