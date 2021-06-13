package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
)

//Get all users
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	var jsonUsers []map[string]string

	//Get all
	database.DB.Find(&users)

	for _, user := range users {
		userItem := map[string]string{}
		userItem["id"] = strconv.Itoa(int(user.Id))
		userItem["login"] = user.Login
		userItem["name"] = user.Name
		userItem["surname"] = user.Surname
		userItem["email"] = user.Email

		jsonUsers = append(jsonUsers, userItem)
	}

	return c.JSON(jsonUsers)
}

//Get user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	//Find by id
	database.DB.Where("id = ?", id).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	var jsonUser = map[string]string{}

	jsonUser["id"] = strconv.Itoa(int(user.Id))
	jsonUser["login"] = user.Login
	jsonUser["name"] = user.Name
	jsonUser["surname"] = user.Surname
	jsonUser["email"] = user.Email


	return c.JSON(jsonUser)
}

//Delete user by id
func DeleteUserById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id := data["id"]


	database.DB.Where("user_id", id).Delete(&models.LogOfUser{})
	database.DB.Delete(&models.User{}, id)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "User with ID:"+id+" deleted",
	})
}

//Update user by id
func UpdateUserById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user := models.User{}
	database.DB.Where("id = ?", data["id"]).First(&user)

	user.Email = data["email"]
	user.Name = data["name"]
	user.Surname = data["surname"]
	user.Login = data["login"]

	pass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user.Password = pass

	database.DB.Save(&user)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "User with ID:"+data["id"]+" updated!",
	})
}

//Get user auth logs
func GetUserLogById(c *fiber.Ctx) error {
	id := c.Params("id")
	logs := []models.LogOfUser{}
	var jsonLog []map[string]string

	database.DB.Where("user_id", id).Find(&logs)

	if (len(logs) != 0) {
		for _,log := range logs {
			logItem := map[string]string{}

			stringId := strconv.Itoa(int(log.Id))
			stringUserID := strconv.Itoa(int(log.UserId))

			logItem["id"] = stringId
			logItem["user_id"] = stringUserID
			logItem["enter_date"] = log.Enter_date.Format("2006-01-02")

			jsonLog = append(jsonLog, logItem)
		}

		c.Status(fiber.StatusOK)
		return c.JSON(jsonLog)
	}

	c.Status(fiber.StatusNotFound)
	return c.JSON(fiber.Map{
		"message": "Logs not found",
	})
}