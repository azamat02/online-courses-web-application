package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

func GetAllPurchasedCourses (c *fiber.Ctx) error {
	var purchasedCourses []models.PurchasedCourses
	var jsonPurchasedCourses []map[string]string

	//Get all
	database.DB.Find(&purchasedCourses)

	for _, purchasedCourse := range purchasedCourses {
		purchasedCourseItem := map[string]string{}
		purchasedCourseItem["id"] = strconv.Itoa(int(purchasedCourse.Id))
		purchasedCourseItem["u_id"] = strconv.Itoa(purchasedCourse.UserId)
		purchasedCourseItem["c_id"] = strconv.Itoa(purchasedCourse.CourseId)
		purchasedCourseItem["purchased_date"] = purchasedCourse.PurchasedDate.Format("2 January 2006 at 15:15")

		jsonPurchasedCourses = append(jsonPurchasedCourses, purchasedCourseItem)
	}

	return c.JSON(jsonPurchasedCourses)
}

func GetAllUserPurchasedCourses (c *fiber.Ctx) error {
	userId := c.Params("id")

	var purchasedCourses []models.PurchasedCourses
	var jsonPurchasedCourses []map[string]string

	//Get all
	database.DB.Where("user_id = ?", userId).Find(&purchasedCourses)

	for _, purchasedCourse := range purchasedCourses {
		purchasedCourseItem := map[string]string{}
		purchasedCourseItem["id"] = strconv.Itoa(int(purchasedCourse.Id))
		purchasedCourseItem["u_id"] = strconv.Itoa(purchasedCourse.UserId)
		purchasedCourseItem["c_id"] = strconv.Itoa(purchasedCourse.CourseId)
		purchasedCourseItem["purchased_date"] = purchasedCourse.PurchasedDate.Format("2 January 2006 at 15:15")

		jsonPurchasedCourses = append(jsonPurchasedCourses, purchasedCourseItem)
	}

	return c.JSON(jsonPurchasedCourses)
}

func CreatePurchasedCourse (c *fiber.Ctx) error {
	//Get data of purchased course
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	purchased_date := time.Now()

	u_id, _ := strconv.Atoi(data["u_id"])
	c_id, _ := strconv.Atoi(data["c_id"])

	//Setting purchased course data
	purchased_course := models.PurchasedCourses{
		UserId: u_id,
		CourseId: c_id,
		PurchasedDate: purchased_date,
	}

	//Creating course row in DB
	database.DB.Create(&purchased_course)

	//Return course
	return c.JSON(purchased_course)
}
