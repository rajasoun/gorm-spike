package db

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig reads the configuration file and loads the database configuration into a map.
func LoadConfig(dbConfigFile string) map[string]string {
	viper.SetConfigFile(dbConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	dbConfig := viper.GetStringMapString("db")
	return dbConfig
}
