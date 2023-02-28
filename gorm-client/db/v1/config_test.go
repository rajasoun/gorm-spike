package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestLoadGormConfig(t *testing.T) {
	// Call the function
	gormConfig, err := LoadGormConfig()
	require.NoError(t, err)

	// Check that the returned config matches the expected values
	assert.Equal(t, false, gormConfig.SkipDefaultTransaction)
	assert.Equal(t, false, gormConfig.FullSaveAssociations)
	assert.Equal(t, false, gormConfig.DryRun)
	assert.Equal(t, false, gormConfig.PrepareStmt)
	assert.Equal(t, false, gormConfig.DisableAutomaticPing)
	assert.Equal(t, false, gormConfig.DisableForeignKeyConstraintWhenMigrating)
	assert.Equal(t, false, gormConfig.IgnoreRelationshipsWhenMigrating)
	assert.Equal(t, false, gormConfig.DisableNestedTransaction)
	assert.Equal(t, false, gormConfig.AllowGlobalUpdate)
	assert.Equal(t, false, gormConfig.QueryFields)
	assert.Equal(t, 0, gormConfig.CreateBatchSize)
}
