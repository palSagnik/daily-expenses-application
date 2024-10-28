package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/daily-expenses-application/database"
)

func GetUserDetails(c *fiber.Ctx) error {
	var userid int
	var err error

	userid_string := c.FormValue("userid")
	if userid_string == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "message": "missing parameters in request"})
	}
	userid, err = strconv.Atoi(userid_string)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "message": "invalid userid"})
	}

	user, err := database.GetUserDetails(c, userid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"success", "message":err.Error()})
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	
	users, err := database.GetUsers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"success", "message":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
