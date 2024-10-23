package database

import "github.com/palSagnik/daily-expenses-application/models"

var (
	users = new(models.User)
	verification = new(models.Verification)
	expense = new(models.Expense)
)


func MigrateUp() error {

	err := DB.AutoMigrate(users, verification, expense)
	if err != nil {
		return err
	}

	return nil
}