package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2/log"
	"github.com/palSagnik/daily-expenses-application/config"
)

var (
	DB *gorm.DB
	err error
)

// connecting to the database
func ConnectDB() error {
	
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT)
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Infof("connected to database: %s\n", config.DB_NAME)
	
	return nil
}

// create database if not present
// func CreateDatabase() error {

// }