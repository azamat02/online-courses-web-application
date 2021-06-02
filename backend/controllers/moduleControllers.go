package controllers

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
)

//Get 1 module
func GetModule(c *fiber.Ctx) error {
	id := c.Params("id")

	var module models.Module

	//Find by id
	database.DB.Where("id = ?", id).First(&module)

	if module.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Module not found",
		})
	}

	var jsonModule = map[string]string{}

	jsonModule["id"] = strconv.Itoa(int(module.Id))
	jsonModule["title"] = module.Title
	jsonModule["c_id"] = strconv.Itoa(module.CourseId)
	jsonModule["number_of_lessons"] = strconv.Itoa(module.Number_of_lessons)

	return c.JSON(jsonModule)
}

//Get modules of course
func GetModulesByCourseId(c *fiber.Ctx) error {
	id := c.Params("id")

	var modules []models.Module
	var jsonModules []map[string]string

	//Find by id
	database.DB.Where("course_id = ?", id).Find(&modules)

	for _, module := range modules {
		moduleItem := map[string]string{}
		moduleItem["id"] = strconv.Itoa(int(module.Id))
		moduleItem["title"] = module.Title
		moduleItem["c_id"] = strconv.Itoa(module.CourseId)
		moduleItem["number_of_lessons"] = strconv.Itoa(module.Number_of_lessons)

		jsonModules = append(jsonModules, moduleItem)
	}

	return c.JSON(jsonModules)
}

//Get all modules
func GetAllModules(c *fiber.Ctx) error {
	var modules []models.Module
	var jsonModules []map[string]string

	//Get all
	database.DB.Find(&modules)

	for _, module := range modules {
		moduleItem := map[string]string{}
		moduleItem["id"] = strconv.Itoa(int(module.Id))
		moduleItem["title"] = module.Title
		moduleItem["c_id"] = strconv.Itoa(module.CourseId)
		moduleItem["number_of_lessons"] = strconv.Itoa(module.Number_of_lessons)

		jsonModules = append(jsonModules, moduleItem)
	}

	return c.JSON(jsonModules)
}

//Create module
func CreateModule(c *fiber.Ctx) error {
	//Get data of module
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	courseId, _ := strconv.Atoi(data["c_id"])
	numberOfLessons,_ := strconv.Atoi(data["number_of_lessons"])

	//Setting module data
	module := models.Module{
		Title: data["title"],
		CourseId: courseId,
		Number_of_lessons: numberOfLessons,
	}

	//Creating module row in DB
	database.DB.Create(&module)

	//Return module
	return c.JSON(module)
}

//Delete module by id
func DeleteModuleById(c *fiber.Ctx) error {
	//Get data of module
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id := data["id"]

	database.DB.Where("module_id", id).Delete(&models.Lesson{})
	database.DB.Where("id", id).Delete(&models.Module{})

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Module with ID:"+id+" deleted",
	})
}

//Update lesson by id
func UpdateModuleById(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	module := models.Module{}
	database.DB.Where("id = ?", data["id"]).First(&module)

	number_of_lessons, _ := strconv.Atoi(data["number_of_lessons"])
	c_id, _ := strconv.Atoi(data["c_id"])

	module.Number_of_lessons = number_of_lessons
	module.Title = data["title"]
	module.CourseId = c_id

	database.DB.Save(&module)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Module with ID:"+data["id"]+" updated!",
	})
}

