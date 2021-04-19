package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "Hospital"
	password = "root"
	sslmode  = "disable"
)

func TestGetConnection(t *testing.T) {
	assert := assert.New(t)
	_, err := GetConnection(host, port, user, dbname, password, sslmode)
	assert.NoError(err, fmt.Errorf("got an error when tried to make connection with database:%w", err))
}

func TestGetConnection2(t *testing.T) {
	assert := assert.New(t)
	password = "roots"
	_, er := GetConnection(host, port, user, dbname, password, sslmode)
	assert.EqualError(er, "got an error when tried to make connection with database:failed to connect to `host=localhost user=postgres database=Hospital`: failed SASL auth (FATAL: password authentication failed for user \"postgres\" (SQLSTATE 28P01))")
}
