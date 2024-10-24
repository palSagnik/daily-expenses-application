package database

import "github.com/palSagnik/daily-expenses-application/models"

var (
	users = new(models.User)
	verification = new(models.Verification)
)


func MigrateUp() error {

	err := DB.AutoMigrate(users, verification)
	if err != nil {
		return err
	}

	return nil
}