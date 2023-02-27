package cmd

import (
	"github.com/rajasoun/gorm-client/db"
	"github.com/rajasoun/gorm-client/log"
)

func Execute() {
	logger := log.NewlogrusLogger()
	logger.Info("DB Stats with Default Connection Pool Settings...")
	handler := db.GetDBHandler()
	handler.PrintDBStats()
	logger.Info("DB Stats with Connection Pool Settings...")
	handler = db.GetDBHandlerWithConnectionPool()
	handler.PrintDBStats()
}
