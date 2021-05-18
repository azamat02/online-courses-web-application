package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

func GetComment(c *fiber.Ctx) error {
	id := c.Params("id")

	var comment models.Comment

	//Find by id
	database.DB.Where("id = ?", id).First(&comment)

	if comment.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Comment not found",
		})
	}

	var jsonComment = map[string]string{}


	jsonComment["id"] = strconv.Itoa(int(comment.Id))
	jsonComment["u_id"] = strconv.Itoa(int(comment.UserId))
	jsonComment["c_id"] = strconv.Itoa(int(comment.CourseId))
	jsonComment["ctext"] = comment.Ctext
	jsonComment["rate"] = strconv.Itoa(int(comment.Rate))
	jsonComment["created_date"] = comment.Created_date.Format("2 January 2006 at 15:04")


	return c.JSON(jsonComment)
}

func GetAllComments(c *fiber.Ctx) error {
	var comments []models.Comment
	var jsonComments []map[string]string

	//Get all
	database.DB.Find(&comments)

	for _, comment := range comments {
		commentItem := map[string]string{}
		commentItem["id"] = strconv.Itoa(int(comment.Id))
		commentItem["u_id"] = strconv.Itoa(int(comment.UserId))
		commentItem["c_id"] = strconv.Itoa(int(comment.CourseId))
		commentItem["ctext"] = comment.Ctext
		commentItem["rate"] = strconv.Itoa(int(comment.Rate))
		commentItem["created_date"] = comment.Created_date.Format("2 January 2006 at 15:04")


		jsonComments = append(jsonComments, commentItem)
	}

	return c.JSON(jsonComments)
}

func GetCommentsByCourseId(c *fiber.Ctx) error {
	courseId := c.Params("id")
	var comments []models.Comment
	var jsonComments []map[string]string

	//Get all
	database.DB.Where("course_id = ?", courseId).Find(&comments)

	for _, comment := range comments {
		commentItem := map[string]string{}
		commentItem["id"] = strconv.Itoa(int(comment.Id))
		commentItem["u_id"] = strconv.Itoa(int(comment.UserId))
		commentItem["c_id"] = strconv.Itoa(int(comment.CourseId))
		commentItem["ctext"] = comment.Ctext
		commentItem["rate"] = strconv.Itoa(int(comment.Rate))
		commentItem["created_date"] = comment.Created_date.Format("2 January 2006 at 15:04")


		jsonComments = append(jsonComments, commentItem)
	}

	return c.JSON(jsonComments)
}

func CreateComment(c *fiber.Ctx) error {
	//Get data of comment
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	created_date := time.Now()
	rate, _ := strconv.Atoi(data["rate"])
	u_id, _ := strconv.Atoi(data["u_id"])
	c_id, _ := strconv.Atoi(data["c_id"])

	//Setting course data
	comment := models.Comment{
		UserId: u_id,
		CourseId: c_id,
		Ctext: data["ctext"],
		Rate: rate,
		Created_date: created_date,
	}

	//Creating comment row in DB
	database.DB.Create(&comment)

	//Return comment
	return c.JSON(comment)
}

