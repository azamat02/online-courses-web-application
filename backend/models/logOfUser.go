package models

import "time"

type LogOfUser struct {
	Id uint `json: "id"`
	UserId int `json: "u_id"`
	Enter_date time.Time `json: "enter_date"`
	User User `json: "-"`
}
