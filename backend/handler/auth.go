package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/daily-expenses-application/models"
	"github.com/palSagnik/daily-expenses-application/utils"
)

func Signup(c *fiber.Ctx) error {
	signup := new(models.User)

	// assigning form values
	signup.Email = c.FormValue("email")
	signup.Name = c.FormValue("name")
	signup.Number = c.FormValue("number")
	signup.Password = c.FormValue("password")
	signup.ConfirmPass = c.FormValue("confirm")

	// handling error if any of the fields are empty
	if signup.Email == "" || signup.Name == ""|| signup.Number == "" || signup.Password == "" || signup.ConfirmPass == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":"all fields must be filled"})
	}

	// removing extra space in form fields
	signup.Email = strings.TrimSpace(signup.Email)
	signup.Password = strings.TrimSpace(signup.Password)
	signup.ConfirmPass = strings.TrimSpace(signup.ConfirmPass)
	signup.Name = strings.TrimSpace(signup.Name)

	// converting email to lowercase
	signup.Email = strings.ToLower(signup.Email)

	// checking whether signup information is valid or not
	isOk, status := utils.VerifySignupInput(signup)
	if !isOk {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":status})
	}

	// TODO:(Check if similar email already exists in the database)

	// Store the hash of the password in the database

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status":"success", "message":"check your email for verification"})
}

// func GetUserDetails(c *fiber.Ctx) error {

// }