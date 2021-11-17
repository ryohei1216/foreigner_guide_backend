package models

import "foreigner_guide/src/database"

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) Create() (err error) {
	db := database.Db
	return db.Create(u).Error
}
