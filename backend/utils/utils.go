package utils

import (
	"fmt"
	"net/mail"

	"github.com/palSagnik/daily-expenses-application/config"
	"github.com/palSagnik/daily-expenses-application/models"
)

func VerifySignupInput(signup *models.User) (bool, string) {
	
	// password verification
	password := signup.Password
	confirmPassword := signup.ConfirmPass

	if len(password) < config.PASS_LEN {
		return false, fmt.Sprintf("password must be atleast of %d characters", config.PASS_LEN)
	}

	if password != confirmPassword {
		return false, "your passwords do not match, please try again"
	}

	// name verification
	name := signup.Name
	if len(name) > config.NAME_LEN {
		return false, fmt.Sprintf("name must be less than %d characters", config.NAME_LEN)
	}

	// email verification
	email := signup.Email
	if len(email) > config.MAIL_LEN {
		return false, fmt.Sprintf("email should not exceed %d characters", config.MAIL_LEN)
	}

	// checking if the email address is valid
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, "not a valid email address"
	}

	return true ,""
}