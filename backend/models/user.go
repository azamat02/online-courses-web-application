package models

type User struct {
	Id uint `json:"id"`
	Login string `json:"login"`
	Password []byte `json:"-"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `gorm:"unique" json:"email"`
}
