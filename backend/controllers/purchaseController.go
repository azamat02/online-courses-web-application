package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"online-courses-app/database"
	"online-courses-app/models"
	"strconv"
	"time"
)

//Get all purchased courses
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

//Get all purchased courses of user
func GetAllUserPurchasedCourses (c *fiber.Ctx) error {
	userId := c.Params("id")

	var purchasedCourses []models.PurchasedCourses
	var jsonPurchasedCourses []map[string]string

	//Get all
	database.DB.Where("user_id = ?", userId).Find(&purchasedCourses)

	for _, purchasedCourse := range purchasedCourses {
		//Get progress of course

		//Get all lessons thar user completed
		completedLessonsLog := []models.CompletedLessonLog{}
		database.DB.Where("user_id = ?", userId).Where("course_id = ?", purchasedCourse.CourseId).Find(&completedLessonsLog)

		//Get all lessons count
		modules := []models.Module{}
		database.DB.Where("course_id = ?", purchasedCourse.CourseId).Find(&modules)

		lessons := []models.Lesson{}
		for _, module := range modules {
			lessonsItem := []models.Lesson{}
			database.DB.Where("module_id = ?", module.Id).Find(&lessonsItem)
			lessons = append(lessons, lessonsItem...)
		}
		lessonsCount := len(lessons)

		progress := float64(0)
		if (lessonsCount != 0) {
			progress = float64((100 * len(completedLessonsLog))) / float64((lessonsCount))
		}

		purchasedCourseItem := map[string]string{}
		purchasedCourseItem["id"] = strconv.Itoa(int(purchasedCourse.Id))
		purchasedCourseItem["u_id"] = strconv.Itoa(purchasedCourse.UserId)
		purchasedCourseItem["c_id"] = strconv.Itoa(purchasedCourse.CourseId)
		purchasedCourseItem["purchased_date"] = purchasedCourse.PurchasedDate.Format("2 January 2006 at 15:15")
		purchasedCourseItem["course_progress"] = strconv.Itoa(int(progress))

		jsonPurchasedCourses = append(jsonPurchasedCourses, purchasedCourseItem)
	}

	return c.JSON(jsonPurchasedCourses)
}

//Create purchase
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

//Check if user purchased course
func CheckIfUserHasPurchasedCourse (c *fiber.Ctx) error {
	courseId := c.Params("id")

	//Get jwt from cookie
	cookie := c.Cookies("jwt")

	//Unparsing token to get user ID
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		//c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	userId, _ := strconv.Atoi(claims.Issuer)
	purchasedCourse := models.PurchasedCourses{}

	database.DB.Where("user_id = ?", userId).Where("course_id = ?", courseId).First(&purchasedCourse)

	if purchasedCourse.Id >0 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "User purchased course",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User not purchased course",
	})
}

func GetUserRecommendations (c *fiber.Ctx) error {
	//Get data
	u_id, _ := strconv.Atoi(c.Params("id"))

	purchasedCourse := models.PurchasedCourses{}

	database.DB.Where("user_id = ?", u_id).First(&purchasedCourse)

	//If user purchased at least one course
	if (purchasedCourse.Id != 0){
		course := models.Course{}
		database.DB.Where("id = ?", purchasedCourse.CourseId).First(&course)

		category := course.Category

		recommendCourses := []models.Course{}
		database.DB.Where("category = ?", category).Where("id != ?", course.Id).Find(&recommendCourses)
		var jsonRecCourses []map[string]string

		for _, course := range recommendCourses {
			courseItem := map[string]string{}
			courseItem["id"] = strconv.Itoa(int(course.Id))
			courseItem["img"] = course.Img
			courseItem["title"] = course.Title
			courseItem["desc"] = course.Description
			courseItem["created_data"] = course.Created_data.Format("2 January 2006")
			courseItem["req"] = course.Req
			courseItem["what_you_will_learn"] = course.What_you_will_learn
			courseItem["category"] = course.Category

			jsonRecCourses = append(jsonRecCourses, courseItem)
		}

		c.Status(fiber.StatusOK)
		return c.JSON(jsonRecCourses)
	}

	//If user hasn't purchased at least one course
	popularCourses := []models.PurchasedCourses{}
	sql := "SELECT course_id,COUNT(course_id) AS count FROM purchased_courses GROUP BY course_id ORDER BY count DESC LIMIT 4"
	database.DB.Raw(sql).Find(&popularCourses)
	var jsonPopCourses []map[string]string

	for _, course := range popularCourses {
		courseId := course.CourseId
		findedCourse := models.Course{}
		database.DB.Where("id = ?", courseId).First(&findedCourse)

		courseItem := map[string]string{}
		courseItem["id"] = strconv.Itoa(int(findedCourse.Id))
		courseItem["img"] = findedCourse.Img
		courseItem["title"] = findedCourse.Title
		courseItem["desc"] = findedCourse.Description
		courseItem["created_data"] = findedCourse.Created_data.Format("2 January 2006")
		courseItem["req"] = findedCourse.Req
		courseItem["what_you_will_learn"] = findedCourse.What_you_will_learn
		courseItem["category"] = findedCourse.Category

		jsonPopCourses = append(jsonPopCourses, courseItem)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(jsonPopCourses)
}
