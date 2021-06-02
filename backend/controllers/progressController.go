package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

//Get user progress in course
func GetCourseProgress(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user_id, _ := strconv.Atoi(data["u_id"])
	course_id, _ := strconv.Atoi(data["c_id"])

	//Get all lessons thar user completed
	completedLessonsLog := []models.CompletedLessonLog{}
	database.DB.Where("user_id = ?", user_id).Where("course_id = ?", course_id).Find(&completedLessonsLog)

	//Get all lessons count
	modules := []models.Module{}
	database.DB.Where("course_id = ?", course_id).Find(&modules)

	lessons := []models.Lesson{}
	for _, module := range modules {
		lessonsItem := []models.Lesson{}
		database.DB.Where("module_id = ?", module.Id).Find(&lessonsItem)
		lessons = append(lessons, lessonsItem...)
	}
	lessonsCount := len(lessons)

	if (lessonsCount != 0) {
		progress := float64((100 * len(completedLessonsLog))) / float64((lessonsCount))
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"progress": progress,
		})
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Lessons for this course not found!",
		})
	}

}

//Get if lesson completed, and complete if not
func CompleteLesson(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user_id, _ := strconv.Atoi(data["u_id"])
	course_id, _ := strconv.Atoi(data["c_id"])
	lesson_id, _ := strconv.Atoi(data["l_id"])

	completedLesson := models.CompletedLessonLog{}

	database.DB.Where("user_id = ?", user_id).Where("course_id = ?", course_id).Where("lesson_id", lesson_id).Find(&completedLesson)

	if (completedLesson.Id != 0) {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "Lesson already completed!",
		})
	}
	completedLesson = models.CompletedLessonLog{
		UserId: user_id,
		CourseId: course_id,
		LessonId: lesson_id,
	}
	log := models.CourseAnalyticsLog{
		CourseId: course_id,
		UserId: user_id,
		Log: "Completed lesson",
		Date: time.Now(),
	}

	database.DB.Create(&completedLesson)
	database.DB.Create(&log)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Lesson completed!",
	})
}

//Check if user completed lesson
func IsLessonCompleted(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user_id, _ := strconv.Atoi(data["u_id"])
	course_id, _ := strconv.Atoi(data["c_id"])
	lesson_id, _ := strconv.Atoi(data["l_id"])

	completedLesson := models.CompletedLessonLog{}

	database.DB.Where("user_id = ?", user_id).Where("course_id = ?", course_id).Where("lesson_id", lesson_id).Find(&completedLesson)

	if (completedLesson.Id != 0) {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "Lesson already completed!",
		})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Lesson not completed!",
	})
}
