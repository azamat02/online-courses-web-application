package models

import "time"

type Comment struct {
	Id uint `json: "id"`
	UserId int `json: "u_id"`
	CourseId int `json: "c_id"`
	Ctext string `json: "ctext"`
	Rate int `json: "rate"`
	Created_date time.Time `json: "create_date"`
	User User `json: "-"`
	Course Course `json: "-"`
}
