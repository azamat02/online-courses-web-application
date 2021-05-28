package models

type CompletedLessonLog struct {
	Id uint `json: "id"`
	UserId int `json: "u_id"`
	CourseId int `json: "c_id"`
	LessonId int `json: "l_id"`
	User User `json: "-"`
	Course Course `json: "-"`
	Lesson Lesson `json: "-"`
}
