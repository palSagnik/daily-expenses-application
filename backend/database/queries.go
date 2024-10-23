package database

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/palSagnik/daily-expenses-application/models"
	"gorm.io/gorm"
)

// user queries
func AddUserToVerify(c *fiber.Ctx, user *models.User) error {
	email := user.Email
	
	// deleting any previous record of user
	result := DB.Delete(&models.Verification{}, "email = ?", email)
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error
	}
	
	// adding user to verification table
	result = DB.Create(&models.Verification{
		Name: user.Name,
		Email: user.Email,
		Number: user.Number,
		Password: user.Password,
	})
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error
	}

	log.Infof("added user with email '%s' for verification", email)
	return nil

}

// deleting user by email -> unique property
func DeleteUser(c *fiber.Ctx) error {
	email := c.Params("email")
	
	result := DB.Delete(&models.User{}, "email = ?", email)
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error
	}

	log.Infof("deleted user with email '%s'", email)
	return nil
}


func AddUser(c *fiber.Ctx, email string) (string, error) {

	// checking if user already exists
	found, err := doesEmailExist(email)
	if found {
		return "user already exists", errors.New("token already verified")
	}
	if err != nil {
		log.Warn(err)
		return err.Error(), err
	}

	// email does not exist hence fetch from verification table
	var verifiedUser models.Verification
	result := DB.Select("name, email, number, password").Where("email = ?", email).First(&verifiedUser)
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error.Error(), result.Error
	}

	// creating user
	result = DB.Create(&models.User{
		Email: verifiedUser.Email,
		Name: verifiedUser.Name,
		Number: verifiedUser.Number,
		Password: verifiedUser.Password,
	})
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error.Error(), result.Error
	}

	// deleting from verification table
	result = DB.Delete(&models.Verification{}, "email = ?", email)
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error.Error(), result.Error
	}
	
	return "", nil
}

// gettting user details
func GetUserDetails (c *fiber.Ctx, userid int) (*models.User, error) {
	var user models.User
	
	log.Infof("fetching user details of userid '%d'", userid)
	result := DB.Select("name, number, email, expense").Where("userid = ?", userid).First(&user)
	if result.Error != nil {
		log.Warn(result.Error)
		return nil, result.Error
	}

	return &user, nil
}


// misc queries
// to validate creds -> during login
func ValidateCreds(c *fiber.Ctx, creds *models.Credentials) error {
	var user models.User

	result := DB.Where("email = ? AND password = ?", creds.Email, creds.Password).First(&user)
	if result.Error != nil {
		log.Warn(result.Error)
		return result.Error
	}

	log.Infof("verified credentials for '%s'", creds.Email)
	return nil
}

// to check whether there is a duplicate email at the time of signup
func doesEmailExist(email string) (bool, error) {

	user := new(models.User)
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		// Error is record not found
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}

		// some other error
		return false, result.Error
	}

	// duplicate email found
	return true, nil
}