package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

//Get 1 course
func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	var course models.Course

	//Find by id
	database.DB.Where("id = ?", id).First(&course)

	if course.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Course not found",
		})
	}

	var jsonCourse = map[string]string{}

	jsonCourse["id"] = strconv.Itoa(int(course.Id))
	jsonCourse["img"] = course.Img
	jsonCourse["title"] = course.Title
	jsonCourse["desc"] = course.Description
	jsonCourse["created_data"] = course.Created_data.Format("2 January 2006")
	jsonCourse["req"] = course.Req
	jsonCourse["what_you_will_learn"] = course.What_you_will_learn

	return c.JSON(jsonCourse)
}

//Get all courses
func GetAllCourses(c *fiber.Ctx) error {
	var courses []models.Course
	var jsonCourses []map[string]string

	//Get all
	database.DB.Find(&courses)

	for _, course := range courses {
		courseItem := map[string]string{}
		courseItem["id"] = strconv.Itoa(int(course.Id))
		courseItem["img"] = course.Img
		courseItem["title"] = course.Title
		courseItem["desc"] = course.Description
		courseItem["created_data"] = course.Created_data.Format("2 January 2006")
		courseItem["req"] = course.Req
		courseItem["what_you_will_learn"] = course.What_you_will_learn

		jsonCourses = append(jsonCourses, courseItem)
	}

	return c.JSON(jsonCourses)
}

//Create course
func CreateCourse(c *fiber.Ctx) error {
	//Get data of course
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	created_date := time.Now()

	//Setting course data
	course := models.Course{
		Img: data["img"],
		Title: data["title"],
		Description: data["desc"],
		Created_data: created_date,
		Req: data["req"],
		What_you_will_learn: data["what_you_will_learn"],
	}

	//Creating course row in DB
	database.DB.Create(&course)

	//Return course
	return c.JSON(course)
}

//Delet course by id
func DeleteCourseById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id := data["id"]

	modules := []models.Module{}

	//Finding and deleting all lessons and modules of course, and also purchases, comments
	database.DB.Where("course_id", id).Find(&modules)
	for _, module := range modules {
		moduleId := module.Id
		database.DB.Where("module_id", moduleId).Delete(&models.Lesson{})
	}
	database.DB.Where("course_id", id).Delete(&models.Module{})
	database.DB.Where("course_id", id).Delete(&models.PurchasedCourses{})
	database.DB.Where("course_id", id).Delete(&models.Comment{})
	database.DB.Delete(&models.Course{}, id)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Course with ID:"+id+" deleted",
	})
}

//Update course by id
func UpdateCourseById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	course := models.Course{}
	database.DB.Where("id = ?", data["id"]).First(&course)

	course.Img = data["img"]
	course.Title = data["title"]
	course.Description = data["desc"]
	course.Req = data["req"]
	course.What_you_will_learn = data["what_you_will_learn"]

	database.DB.Save(&course)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Course with ID:"+data["id"]+" updated!",
	})
}

//Get course rating by comments
func GetCourseRating(c *fiber.Ctx) error {
	id := c.Params("id")

	comments := []models.Comment{}

	//Getting all comments of course
	database.DB.Where("course_id", id).Find(&comments)

	if (len(comments) != 0) {
		total_rates_count := 0
		sum_rates := 0

		rate_0_count := 0
		rate_1_count := 0
		rate_2_count := 0
		rate_3_count := 0
		rate_4_count := 0
		rate_5_count := 0

		for _,comment := range comments {
			rating := comment.Rate
			total_rates_count ++
			sum_rates += rating

			if (rating == 5) {
				rate_5_count ++
			}
			if (rating == 4) {
				rate_4_count ++
			}
			if (rating == 3) {
				rate_3_count ++
			}
			if (rating == 2) {
				rate_2_count ++
			}
			if (rating == 1) {
				rate_1_count ++
			}
			if (rating == 0) {
				rate_0_count ++
			}
		}

		total_rating := float64(((5*rate_5_count)+(4*rate_4_count)+(3*rate_3_count)+(2*rate_2_count)+(1*rate_1_count))) / float64((total_rates_count))
		rating_5 := (100 * rate_5_count) / total_rates_count
		rating_4 := (100 * rate_4_count) / total_rates_count
		rating_3 := (100 * rate_3_count) / total_rates_count
		rating_2 := (100 * rate_2_count) / total_rates_count
		rating_1 := (100 * rate_1_count) / total_rates_count
		rating_0 := (100 * rate_0_count) / total_rates_count

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"total_rating": total_rating,
			"rating_5": rating_5,
			"rating_4": rating_4,
			"rating_3": rating_3,
			"rating_2": rating_2,
			"rating_1": rating_1,
			"rating_0": rating_0,
		})
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Comments for this course not found",
		})
	}
}