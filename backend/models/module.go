package models

type Module struct {
	Id uint `json: "id"`
	Title string `json: "title"`
	CourseId int `json: "c_id"`
	Number_of_lessons int `json: "number_of_lessons"`
	Course Course `json: "-"`
}
