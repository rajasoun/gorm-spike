package db

import (
	"database/sql"

	"github.com/rajasoun/gorm-client/log"
)

// GetDBStats returns a DBStats struct containing the database connection statistics.
func (h *DBHandler) GetDBStats() sql.DBStats {
	return h.DB.Stats()
}

// PrintDBStats prints the database connection statistics to the console.
func (h *DBHandler) PrintDBStats() {
	stats := h.GetDBStats()
	logger := log.NewZeroLogger()
	logger.Printf("Max Open Connections: %d", stats.MaxOpenConnections)
	logger.Printf("    Open Connections: %d", stats.OpenConnections)
	logger.Printf("  In Use Connections: %d", stats.InUse)
	logger.Printf("    Idle Connections: %d", stats.Idle)
}
