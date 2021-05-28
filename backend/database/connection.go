package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"online-courses-app/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Azamat2341! dbname=online_courses port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Course{})
	connection.AutoMigrate(&models.Module{})
	connection.AutoMigrate(&models.Lesson{})
	connection.AutoMigrate(&models.Comment{})
	connection.AutoMigrate(&models.LogOfUser{})
	connection.AutoMigrate(&models.PurchasedCourses{})
	connection.AutoMigrate(&models.CompletedLessonLog{})
	connection.AutoMigrate(&models.CourseAnalyticsLog{})
}
