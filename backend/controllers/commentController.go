package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"strings"
	"time"
)

//Get 1 comment by id
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

//Get all comments
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

//Get comments of course
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

//Creating comment adn rate by comment text
func CreateComment(c *fiber.Ctx) error {
	//Get data of comment
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	created_date := time.Now()
	rate := 0
	u_id, _ := strconv.Atoi(data["u_id"])
	c_id, _ := strconv.Atoi(data["c_id"])
	feedback := data["ctext"]


	synonimsForGood := []string{"good", "great", "excellent",
									"pleasant", "amazing", "fantastic",
									"valuable", "perfect", "super",
									"nice", "fine", "satisfying",
									"brilliant", "awesome", "wonderful",
									"first-rate", "superior", "the best", "very good", "best"}
	synonimsForBad := []string{"bad", "sad", "awful", "lousy", "worse", "dreadful", "negative", "very bad", "uncool"}

	badCount := 0
	goodCount := 0

	//Count "bad words"
	for _, word := range synonimsForGood {
		if (strings.Contains(feedback, string(word))) {
			goodCount++
		}
	}

	//Count "good words"
	for _, word := range synonimsForBad {
		if (strings.Contains(feedback, word)) {
			badCount++
		}
	}

	total := badCount + goodCount
	if (total != 0) {
		percentageOfGood := (100*goodCount) / total

		if (percentageOfGood>=80 && percentageOfGood <= 100) {
			rate = 5
		}
		if (percentageOfGood >= 60 && percentageOfGood < 80) {
			rate = 4
		}
		if (percentageOfGood >= 40 && percentageOfGood < 60) {
			rate = 3
		}
		if (percentageOfGood >= 20 && percentageOfGood < 40) {
			rate = 2
		}
		if (percentageOfGood >= 10 && percentageOfGood < 20) {
			rate = 1
		}
	}

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

//Check if user leaved feadback to the course
func Check(c *fiber.Ctx) error {
	//Get data of comment
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user_id := data["u_id"]
	course_id := data["c_id"]

	comment := models.Comment{}

	//Find if user leave comment
	database.DB.Where("course_id = ?", course_id).Where("user_id = ?", user_id).First(&comment)

	if (comment.UserId != 0) {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "User rated",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User not rated",
	})
}

