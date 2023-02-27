package db

import (
	"database/sql"
	"log"

	"gorm.io/gorm"
)

type DBHandler struct {
	DB *sql.DB
}

// Get the sql.DB object from the gorm.DB object.
func getDbClient(db *gorm.DB) *sql.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB object: %v", err)
	}
	return sqlDB
}

// GetDBHandler returns a pointer to a DBHandler struct.
func GetDBHandler() *DBHandler {
	con := GetConnection()
	dbClient := getDbClient(con)
	return &DBHandler{DB: dbClient}
}

// GetDBHandler returns a pointer to a DBHandler struct with connection pool.
func GetDBHandlerWithConnectionPool() *DBHandler {
	con := GetConnectionWithConnectionPoll()
	dbClient := getDbClient(con)
	return &DBHandler{DB: dbClient}
}
