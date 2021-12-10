package models

import (
	"foreigner_guide/src/database"
)
type User struct {
	Id        string `json:"id"` 
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	Image     string `json:image`
}

func (u *User) Create() (err error) {
	db := database.Db
	return db.Create(u).Error
}

func GetAll() (users []User ,err error) {
	db := database.Db
	users = []User{}
	err = db.Find(&users).Error
	return users, err
}

func (u *User) GetBySignIn() (user User, err error) {
	db := database.Db
	user = User{}
	err = db.Where("email = ? AND password = ?", u.Email, u.Password).Find(&user).Error
	return user, err
}

func (u *User) GetByArea() (users []User, err error) {
	db := database.Db
	users = []User{}
	err = db.Where("area = ?", u.Area).Find(&users).Error
	return users, err
}
