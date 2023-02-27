package test

import (
	"reflect"
	"testing"

	"github.com/rajasoun/gorm-client/db/v0"
	"github.com/rajasoun/gorm-client/test/load"
)

func TestLoadConfig(t *testing.T) {
	expected := map[string]string{
		"dbname":   "gorm_spike",
		"user":     "root",
		"password": "example",
		"host":     "127.0.0.1",
		"port":     "3306",
		"charset":  "utf8mb4",
		"loc":      "Local",
	}
	actual := db.LoadConfig("db.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("LoadConfig() returned %+v, expected %+v", actual, expected)
	}
}

func TestExecuteWithConnectionPool(t *testing.T) {
	handler := db.GetDBHandlerWithConnectionPool()
	stats := handler.GetDBStats()

	testCases := []struct {
		name           string
		actual         int
		expected       int
		errorMsgFormat string
	}{
		{"Max open connections", stats.MaxOpenConnections, 100, "Max open connections = %d; expected 100"},
		{"Open connections", stats.OpenConnections, 1, "Open connections = %d; expected 1"},
		{"In-use connections", stats.InUse, 0, "In-use connections = %d; expected 0"},
		{"Idle connections", stats.Idle, 1, "Idle connections = %d; expected 1"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf(tc.errorMsgFormat, tc.actual)
			}
		})
	}
}

func TestExecuteWithoutConnectionPool(t *testing.T) {
	handler := db.GetDBHandler()
	stats := handler.GetDBStats()

	testCases := []struct {
		name           string
		actual         int
		expected       int
		errorMsgFormat string
	}{
		{"Max open connections", stats.MaxOpenConnections, 0, "Max open connections = %d; expected 0"},
		{"Open connections", stats.OpenConnections, 1, "Open connections = %d; expected 1"},
		{"In-use connections", stats.InUse, 0, "In-use connections = %d; expected 0"},
		{"Idle connections", stats.Idle, 1, "Idle connections = %d; expected 1"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf(tc.errorMsgFormat, tc.actual)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	repo, err := load.NewUserRepository()
	if err != nil {
		t.Fatal("Failed to create user repository", err)
	}
	// Generate a random user
	err = repo.NewUser()
	if err != nil {
		t.Fatal("Failed to create user", err)
	}
}

func TestCreateUsersWithDefaultConnectionPool(t *testing.T) {
	repo, err := load.NewUserRepository()
	if err != nil {
		t.Fatal("Failed to create user repository", err)
	}
	// Create the users using multiple goroutines
	numOfGoroutines := 100
	numUsersToCreatePerGoroutine := 500

	repo.CreateUsers(numOfGoroutines, numUsersToCreatePerGoroutine)
	t.Logf("Created all users")

	// Print the database connection statistics
	stats := db.GetDBHandler().GetDBStats()
	//stats := db.GetDBHandlerWithConnectionPoll().GetDBStats()
	t.Logf("Max Open Connections: %d", stats.MaxOpenConnections)
	t.Logf("    Open Connections: %d", stats.OpenConnections)
	t.Logf("  In Use Connections: %d", stats.InUse)
	t.Logf("    Idle Connections: %d", stats.Idle)

	// Get Total number of users
	totalUsers := repo.GetTotalUsers()
	t.Logf("Total number of users: %d", totalUsers)
}

func TestCreateUsersWithConnectionPoolConfigured(t *testing.T) {
	repo, err := load.NewUserRepositoryWithConnectionPool()
	if err != nil {
		t.Fatal("Failed to create user repository", err)
	}
	// Create the users using multiple goroutines
	numOfGoroutines := 100
	numUsersToCreatePerGoroutine := 5000

	repo.CreateUsers(numOfGoroutines, numUsersToCreatePerGoroutine)
	t.Logf("Created all users")

	// Print the database connection statistics
	stats := db.GetDBHandler().GetDBStats()
	//stats := db.GetDBHandlerWithConnectionPoll().GetDBStats()
	t.Logf("Max Open Connections: %d", stats.MaxOpenConnections)
	t.Logf("    Open Connections: %d", stats.OpenConnections)
	t.Logf("  In Use Connections: %d", stats.InUse)
	t.Logf("    Idle Connections: %d", stats.Idle)

	// Get Total number of users
	totalUsers := repo.GetTotalUsers()
	t.Logf("Total number of users: %d", totalUsers)
}
