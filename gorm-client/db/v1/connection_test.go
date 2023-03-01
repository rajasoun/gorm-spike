package db

import (
	"log"
	"testing"

	"github.com/rajasoun/gorm-client/db/v1/test"
	"github.com/stretchr/testify/assert"
)

func TestNewDBWithDockerTest(t *testing.T) {
	pool, resource := test.InitTestDocker("3306")
	// stop container after test
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Println("Failed to purge resource")
		}
		log.Println("Purged resource")
	}()

	dbConn, err := NewDB()

	assert.NoError(t, err, "Failed to connect to database")
	assert.NotNil(t, dbConn, "Database connection is nil")
}

func TestNewDBWithInMemoryDB(t *testing.T) {
	pool, resource := test.InitTestDocker("3306")
	// stop container after test
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Println("Failed to purge resource")
		}
		log.Println("Purged resource")
	}()

	dbConn, err := NewDB()

	assert.NoError(t, err, "Failed to connect to database")
	assert.NotNil(t, dbConn, "Database connection is nil")
}
