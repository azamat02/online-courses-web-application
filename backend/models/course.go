package models

import "time"

type Course struct {
	Id uint `json:"id"`
	Img string `json:"img"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Created_data time.Time `json:"created_data"`
	Req string `json:"req"`
	What_you_will_learn string `json:"what_you_will_learn"`
	Category string `json:"category"`
}
