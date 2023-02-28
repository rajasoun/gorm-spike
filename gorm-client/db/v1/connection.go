package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (db *DB) GetConnection() *gorm.DB {
	return db.DB
}

func (db *DB) CloseConnection() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func NewDB() (*gorm.DB, error) {
	config, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)

	gormConfig, err := LoadGormConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)

	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %s", err)
	}
	return db, nil
}
