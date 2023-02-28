package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	// Initialize a new database connection.
	dbConn, err := NewDB()

	assert.NoError(t, err, "Failed to connect to database")
	assert.NotNil(t, dbConn, "Database connection is nil")
}
