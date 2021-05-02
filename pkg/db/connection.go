package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)
	connection, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("got an error when tried to make connection with database:%w", err)
	}
	return connection, nil
}
