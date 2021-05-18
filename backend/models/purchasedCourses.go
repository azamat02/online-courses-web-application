package models

import "time"

type PurchasedCourses struct {
	Id uint `json: "id"`
	UserId int `json: "u_id"`
	CourseId int `json: "c_id"`
	PurchasedDate time.Time `json: "purchased_date"`
	User User `json: "-"`
	Course Course `json: "-"`
}