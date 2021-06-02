package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
)

//Get 1 lesson by id
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

//Get module lessons
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

//Get all lessons
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

//Create lesson
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

//Delete lesson by id
func DeleteLessonById(c *fiber.Ctx) error {
	//Get data of lesson
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id := data["id"]

	database.DB.Where("id", id).Delete(&models.Lesson{})

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Lesson with ID:"+id+" deleted",
	})
}

//Update lesson by id
func UpdateLessonById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	lesson := models.Lesson{}
	database.DB.Where("id = ?", data["id"]).First(&lesson)

	m_id, _ := strconv.Atoi(data["m_id"])

	lesson.Type = data["type"]
	lesson.ModuleId = m_id
	lesson.Title = data["title"]
	lesson.Link = data["link"]
	lesson.Content = data["content"]

	database.DB.Save(&lesson)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Lesson with ID:"+data["id"]+" updated!",
	})
}

func GetNextLesson(c *fiber.Ctx)  error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	courseId := data["c_id"]
	lessonId, _ := strconv.Atoi(data["l_id"])
	modules := []models.Module{}
	lessons := []models.Lesson{}

	database.DB.Where("course_id = ?", courseId).Find(&modules)

	for _, module:= range modules {
		l := []models.Lesson{}
		database.DB.Where("module_id", module.Id).Find(&l)
		lessons = append(lessons, l...)
	}

	nextLessonId := 0

	for _, lesson:= range lessons {
		if (lessonId+1 == int(lesson.Id)) {
			nextLessonId = int(lesson.Id)
		}
	}

	if (nextLessonId != 0) {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"nextLessonId": nextLessonId,
		})
	}

	c.Status(fiber.StatusNotFound)
	return c.JSON(fiber.Map{
		"message": "Next lesson not found",
	})
}

