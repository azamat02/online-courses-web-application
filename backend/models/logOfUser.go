package models

import "time"

type LogOfUser struct {
	Id uint
	UserId int
	Enter_date time.Time
	User User
}
