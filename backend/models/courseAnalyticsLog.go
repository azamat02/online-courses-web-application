package models

import "time"

type CourseAnalyticsLog struct {
	Id uint `json: "id"`
	UserId int `json: "u_id"`
	CourseId int `json: "c_id"`
	Log string `json: "log"`
	Date time.Time
	User User `json: "-"`
	Course Course `json: "-"`
}

