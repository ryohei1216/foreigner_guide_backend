package models

import (
	"fmt"
	"foreigner_guide/src/database"
	"log"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	Area      string `json:"area"`
}

func (u *User) Create() (err error) {
	db := database.Db
	type applyUser struct {
		UserId  string
		GuideId string
		Status  string
	}
	//applyUserListのDBも作成
	err = db.Table("apply_user_list" + u.Id).AutoMigrate(&applyUser{})
	if err != nil {
		log.Println(err)
	}
	//appliedUserListのDBも作成
	err = db.Table("applied_user_list" + u.Id).AutoMigrate(&applyUser{})
	if err != nil {
		log.Println(err)
	}

	//chatroomのDB作成
	err = db.Table("messages_"+u.Id).AutoMigrate(&Message{})
	if err != nil {
		log.Println(err)
	}
	return db.Create(u).Error
}

func GetAll() (users []User, err error) {
	db := database.Db
	users = []User{}
	err = db.Find(&users).Error
	return users, err
}

func (u *User) GetById() (user User, err error) {
	db := database.Db
	user = User{}
	err = db.Where("id = ?", u.Id).Find(&user).Error
	return user, err
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

func (u *User) GetByApply() (users []User, err error) {
	db := database.Db
	type applyUser struct {
		UserId  string
		GuideId string
		Status  string
	}
	userList := []applyUser{}
	fmt.Println(u.Id)
	err = db.Table("applied_user_list" + u.Id).Find(&userList).Error
	if err != nil {
		log.Println(err)
	}

	var applyUsersid []string
	for _, v := range userList {
		applyUsersid = append(applyUsersid, v.UserId)
	}

	users = []User{}
	for _, v := range applyUsersid {
		user := User{}
		fmt.Println(v)
		err = db.Where("id = ?", v).Find(&user).Error
		users = append(users, user)
	}
	fmt.Println(userList)
	fmt.Println(applyUsersid)
	fmt.Println(users)

	return users, err
}