package db

import (
	"testing"

	"github.com/kyma-incubator/reconciler/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	t.Run("Do not block invalid queries", func(t *testing.T) {
		blockQueries := false
		validator := NewValidator(blockQueries, logger.NewOptionalLogger(true))
		query := "invalid query"
		err := validator.Validate(query)
		require.NoError(t, err)
	})
	t.Run("Block invalid queries", func(t *testing.T) {
		blockQueries := true
		validator := NewValidator(blockQueries, logger.NewOptionalLogger(true))
		query := "invalid query"
		err := validator.Validate(query)
		require.Error(t, err)
	})

	//Create validator for the rest of the tests, which blocks invalid queries
	validator := NewValidator(true, logger.NewOptionalLogger(true))
	t.Run("Validate valid insert query", func(t *testing.T) {
		query := "INSERT INTO mockTable (col_1, col_3) VALUES ($1, $2) RETURNING col_1, col_2, col_3"
		err := validator.Validate(query)
		require.NoError(t, err)
	})
	t.Run("Validate invalid insert query", func(t *testing.T) {
		query := "INSERT INTO mockTable (col_1, col_3) VALUES (val_1, val_2) RETURNING col_1, col_2, col_3"
		err := validator.Validate(query)
		require.Error(t, err)
	})
	t.Run("Validate valid select query", func(t *testing.T) {
		query := "SELECT col FROM table WHERE y=$1"
		err := validator.Validate(query)
		require.NoError(t, err)
	})
	t.Run("Validate invalid select query", func(t *testing.T) {
		query := "SELECT col FROM table WHERE y=x"
		err := validator.Validate(query)
		require.Error(t, err)
	})
	t.Run("Validate valid delete query", func(t *testing.T) {
		query := "DELETE FROM mockTable WHERE col_1=$1 AND col_2=$2"
		err := validator.Validate(query)
		require.NoError(t, err)
	})
	t.Run("Validate invalid delete query", func(t *testing.T) {
		query := "DELETE FROM mockTable WHERE col_1=v1 AND col_2=v2"
		err := validator.Validate(query)
		require.Error(t, err)
	})
}
