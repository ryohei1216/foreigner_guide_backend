package models

import "foreigner_guide/src/database"

type User struct {
	Id        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string	`json:"lastName"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}

func (u *User) Create() (err error) {
	db := database.Db
	return db.Create(u).Error
}


func ( u *User) GetUser() (user User, err error) {
	db := database.Db
	err = db.Where("email = ? AND password = ?", u.Email, u.Password).Find(&user).Error
	return user, err
}
