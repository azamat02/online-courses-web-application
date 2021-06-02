package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

//Get course analytics of user by id
func GetCourseAnalyticsByUserId(c *fiber.Ctx) error {
	//Get data
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	logs := []models.CourseAnalyticsLog{}
	var jsonLogs []map[string]string

	user_id := data["u_id"]
	course_id := data["c_id"]

	database.DB.Where("user_id = ?",user_id).Where("course_id = ?", course_id).Find(&logs)

	for _, log := range logs {
		logItem := map[string]string{}
		logItem["log"] = log.Log
		logItem["Date"] = log.Date.Format("02 January 2006 at 15:04")

		jsonLogs = append(jsonLogs, logItem)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(jsonLogs)
}

//Create analytics log when user completed lesson or entered to it
func CreateCourseAnalyticsLog(c *fiber.Ctx) error {
	//Get data
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	course_id, _ := strconv.Atoi(data["c_id"])
	user_id, _ := strconv.Atoi(data["u_id"])

	log := models.CourseAnalyticsLog{
		CourseId: course_id,
		UserId: user_id,
		Log: "Entered to lesson",
		Date: time.Now(),
	}

	database.DB.Create(&log)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}
