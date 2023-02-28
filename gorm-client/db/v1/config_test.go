package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("load config", func(t *testing.T) {
		config, err := LoadDBConfig()
		assert.NoError(t, err)
		assert.Equal(t, "localhost", config.Host)
		assert.Equal(t, 3306, config.Port)
		assert.Equal(t, "dbuser", config.User)
		assert.Equal(t, "dbpassword", config.Password)
		assert.Equal(t, "dbname", config.DBName)
		assert.Equal(t, 5, config.MaxIdleConns)
		assert.Equal(t, 20, config.MaxOpenConns)
		assert.Equal(t, 10, config.ConnMaxLifetime)
	})
}
