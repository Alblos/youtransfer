package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var connection *gorm.DB = nil

func Connect() {
	if connection == nil {
		dsn := buildDsn()
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		connection = db
	}

	// Ping the database to check if the connection is successful
	tx := connection.Exec("SELECT 1")
	if tx.Error != nil {
		log.Fatalf("Error connecting to database: %v", tx.Error)
	} else {
		log.Println("Connected successfully to database")
	}
}

func GetConnection() *gorm.DB {
	if connection == nil {
		Connect()
	}
	return connection
}

func buildDsn() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatalf("DB_HOST is not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatalf("DB_USER is not set")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatalf("DB_PASSWORD is not set")
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatalf("DB_NAME is not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatalf("DB_PORT is not set")
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return "host=" + "localhost" + " user=" + "postgres" + " password=" + "password" + " dbname=" + "postgres" + " port=" + "5432" + " sslmode=disable TimeZone=Asia/Shanghai"
}
