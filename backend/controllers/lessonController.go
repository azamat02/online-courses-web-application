package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
)

func GetLesson(c *fiber.Ctx) error {
	id := c.Params("id")

	var lesson models.Lesson

	//Find by id
	database.DB.Where("id = ?", id).First(&lesson)

	if lesson.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Lesson not found",
		})
	}

	var jsonLesson = map[string]string{}

	jsonLesson["id"] = strconv.Itoa(int(lesson.Id))
	jsonLesson["type"] = lesson.Type
	jsonLesson["m_id"] = strconv.Itoa(lesson.ModuleId)
	jsonLesson["title"] = lesson.Title
	jsonLesson["link"] = lesson.Link
	jsonLesson["content"] = lesson.Content

	return c.JSON(jsonLesson)
}

func GetLessonsByModuleId(c *fiber.Ctx) error {
	id := c.Params("id")

	var lessons []models.Lesson
	var jsonLessons []map[string]string

	//Find by id
	database.DB.Where("module_id = ?", id).Find(&lessons)

	for _, lesson := range lessons {
		lessonItem := map[string]string{}
		lessonItem["id"] = strconv.Itoa(int(lesson.Id))
		lessonItem["type"] = lesson.Type
		lessonItem["m_id"] = strconv.Itoa(lesson.ModuleId)
		lessonItem["title"] = lesson.Title
		lessonItem["link"] = lesson.Link
		lessonItem["content"] = lesson.Content

		jsonLessons = append(jsonLessons, lessonItem)
	}

	return c.JSON(jsonLessons)
}

func GetAllLessons(c *fiber.Ctx) error {
	var lessons []models.Lesson
	var jsonLessons []map[string]string

	//Get all
	database.DB.Find(&lessons)

	for _, lesson := range lessons {
		lessonItem := map[string]string{}
		lessonItem["id"] = strconv.Itoa(int(lesson.Id))
		lessonItem["type"] = lesson.Type
		lessonItem["m_id"] = strconv.Itoa(lesson.ModuleId)
		lessonItem["title"] = lesson.Title
		lessonItem["link"] = lesson.Link
		lessonItem["content"] = lesson.Content

		jsonLessons = append(jsonLessons, lessonItem)
	}

	return c.JSON(jsonLessons)
}

func CreateLesson(c *fiber.Ctx) error {
	//Get data of lesson
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	moduleId, _ := strconv.Atoi(data["m_id"])

	//Setting lesson data
	lesson := models.Lesson{
		Type: data["type"],
		ModuleId: moduleId,
		Title: data["title"],
		Link: data["link"],
		Content: data["content"],
	}

	//Creating lesson row in DB
	database.DB.Create(&lesson)

	//Return module
	return c.JSON(lesson)
}


