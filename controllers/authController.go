// Auth Controller
// Register/Login and all that good stuff.
package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/scopehs/tutorial/database"
	"github.com/scopehs/tutorial/models"
	"github.com/scopehs/tutorial/util"
)

// Register Function
func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)

}

// Login Yeoooo
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})

	}

	// Generate a JWT Token based on User ID
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	// Get the Cookie!
	cookie := c.Cookies("jwt")

	// Parses the JWT/Cookie
	id, _ := util.ParseJWT(cookie)

	// Declare a varaiable as User struct
	var user models.User

	// Create a database query to find the user and point the result into user
	database.DB.Where("id = ?", id).First(&user)

	// Return user
	return c.JSON(user)
}

// Logout, by adding a new cookie with no value and -1 Hour from now.
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Get the Cookie!
	cookie := c.Cookies("jwt")

	// Parses the JWT/Cookie
	id, _ := util.ParseJWT(cookie)

	userId, _ := strconv.Atoi(id)
	// Declare a varaiable as User struct
	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	// Updates the user
	database.DB.Where(&user).Where("id = ?", id).Updates(data)

	// Return user
	return c.JSON(user)

}

func UpdatePassword(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	// Get the Cookie!
	cookie := c.Cookies("jwt")

	// Parses the JWT/Cookie
	id, _ := util.ParseJWT(cookie)

	userId, _ := strconv.Atoi(id)

	// Declare a varaiable as User struct
	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])

	// Updates the user
	database.DB.Model(&user).Updates(user)

	// Return user
	return c.JSON(user)

}
