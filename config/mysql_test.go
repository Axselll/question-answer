package config

import (
	"os"
	"qa/exception"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	err := godotenv.Load("../.env")
	exception.CheckError(err)

	db := NewDB(os.Getenv("ENGINE"), os.Getenv("ADDRESS_DB"))
	db.SetMaxIdleConns(0)

	t.Run("Ping db", func(t *testing.T) {
		if db.Ping() != nil {
			assert.Error(t, assert.AnError)
		}
	})
	db.Close()

	t.Run("Check conn", func(t *testing.T) {
		if db.Stats().OpenConnections != 0 {
			assert.Error(t, assert.AnError)
		}
	})
}
