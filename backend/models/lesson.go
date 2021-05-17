package models

type Lesson struct {
	Id uint `json:"id"`
	Type string `json:"type"`
	ModuleId int `json:"m_id"`
	Title string `json:"title"`
	Link string `json:"link"`
	Content string `json:"content"`
	Module Module `json: "-"`
}