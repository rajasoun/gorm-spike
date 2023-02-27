package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Build the DSN string from the database configuration file db.yml
func buildDSN() string {
	dbConfig := LoadConfig("db.yml")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["dbname"],
		dbConfig["charset"],
		url.QueryEscape(dbConfig["loc"]),
	)
	return dsn
}

func CreateDatabase() {
	db, err := sql.Open("mysql", buildDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.Close()
	dbConfig := LoadConfig("db.yml")
	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbConfig["dbname"])
}

type DB struct {
	DB *gorm.DB
}

// Open a connection to the MySQL database.
func GetConnection() *gorm.DB {
	dsn := buildDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

// Open a connection to the MySQL database with connection pool.
func GetConnectionWithConnectionPoll() *gorm.DB {
	dsn := buildDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB object: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
