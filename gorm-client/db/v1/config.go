package db

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
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

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

// LoadConfig loads config from config.yaml
func LoadDBConfig() (*Config, error) {
	initConfig()
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

// GetGormConfig returns gorm config
func LoadGormConfig() (*gorm.Config, error) {
	initConfig()
	err := viper.ReadInConfig() // read the config file
	if err != nil {
		return nil, err
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction:                   viper.GetBool("gorm.skipDefaultTransaction"),
		FullSaveAssociations:                     viper.GetBool("gorm.fullSaveAssociations"),
		DryRun:                                   viper.GetBool("gorm.dryRun"),
		PrepareStmt:                              viper.GetBool("gorm.prepareStmt"),
		DisableAutomaticPing:                     viper.GetBool("gorm.disableAutomaticPing"),
		DisableForeignKeyConstraintWhenMigrating: viper.GetBool("gorm.disableForeignKeyConstraintWhenMigrating"),
		IgnoreRelationshipsWhenMigrating:         viper.GetBool("gorm.ignoreRelationshipsWhenMigrating"),
		DisableNestedTransaction:                 viper.GetBool("gorm.disableNestedTransaction"),
		AllowGlobalUpdate:                        viper.GetBool("gorm.allowGlobalUpdate"),
		QueryFields:                              viper.GetBool("gorm.queryFields"),
		CreateBatchSize:                          viper.GetInt("gorm.createBatchSize"),
	}
	return gormConfig, nil
}
