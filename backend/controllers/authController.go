package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

const SecretKey = "secret"

//Registration
func SignUp(c *fiber.Ctx) error{
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	email := data["email"]

	userToCheck := models.User{}

	database.DB.Where("email = ?", email).First(&userToCheck)

	if (userToCheck.Id != 0) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with current email already exist",
		})
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	//Setting user data
	user := models.User{
		Name: data["name"],
		Surname: data["surname"],
		Login: data["login"],
		Email: data["email"],
		Password: pass,
	}

	//Creating user row in DB
	database.DB.Create(&user)

	//Return user
	return c.JSON(user)
}

//Login
func SignIn(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}


	var user models.User

	//Get user by email
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	log := models.LogOfUser{
		UserId: int(user.Id),
		Enter_date: time.Now(),
	}

	//Creating log of entered user
	database.DB.Create(&log)

	//Creating jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour*24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not sign in",
		})
	}

	//Creating cookie to store jwt token
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour*23),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

//Get authorized user info by jwt token
func User(c *fiber.Ctx) error {
	//Get jwt from cookie
	cookie := c.Cookies("jwt")

	//Unparsing token to get user ID
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		//c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	//Find user by ID which we get from jwt token
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

//Logout
func Logout(c *fiber.Ctx) error {
	//Destruct cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return (c.JSON(fiber.Map{
		"message": "Success",
	}))
}

//Check if user is admin
func IsAdmin(c *fiber.Ctx) error {
	//Get jwt from cookie
	cookie := c.Cookies("jwt")

	//Unparsing token to get user ID
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		//c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(claims.Issuer)

	if id == 1 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "User is admin",
		})
	}

	c.Status(fiber.StatusNotFound)
	return c.JSON(fiber.Map{
		"message": "User is not admin",
	})
}
