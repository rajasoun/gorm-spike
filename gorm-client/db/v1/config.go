package db

import (
	"github.com/spf13/viper"
)

// Config is a struct for database config
type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

// LoadConfig loads config from config.yaml
func LoadDBConfig() (*Config, error) {
	viper.SetConfigName("config") // remove file extension
	viper.SetConfigType("yaml")   // set config type explicitly
	viper.AddConfigPath(".")      // add current directory as search path
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		Host:            viper.GetString("database.host"),
		Port:            viper.GetInt("database.port"),
		User:            viper.GetString("database.user"),
		Password:        viper.GetString("database.password"),
		DBName:          viper.GetString("database.dbname"),
		MaxIdleConns:    viper.GetInt("database.max_idle_conns"),
		MaxOpenConns:    viper.GetInt("database.max_open_conns"),
		ConnMaxLifetime: viper.GetInt("database.conn_max_lifetime"),
	}

	return config, nil
}
