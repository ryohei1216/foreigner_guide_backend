package controllers

import (
	"foreigner_guide/src/database"
	"foreigner_guide/src/models"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = database.Db


func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func CreateUser(c *gin.Context) {
	var user models.User
	user.Id = RandomString(5)
	c.BindJSON(&user)

	err := user.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"create_user": "failed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"create_user": "success",
	})

}

func GetAllUsers(c *gin.Context) {
	users, err := models.GetAll()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetSignInUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
	}
	getUser, err := user.GetBySignIn()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"getUser": getUser,
	})
}

func GetUsersByArea(c *gin.Context) {
	area := c.Query("area")
	modelUser := models.User{
		Area: area,
	}
	users, err := modelUser.GetByArea()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func SaveApplyUser (c *gin.Context) {
	type ApplyUser struct {
		UserId 	string
		GuideId string
		Status  string
	}
	applyUser := ApplyUser{}
	c.BindJSON(&applyUser)
	//userの申請ユーザーリストDBに挿入
	db.Table("apply_user_list"+applyUser.UserId).Create(&applyUser)
	db.Table("applied_user_list"+applyUser.GuideId).Create(&applyUser)
}

func GetApplyUsers (c *gin.Context) {
	id := c.Query("id")
	u := models.User{Id: id}
	users, err := u.GetByApply()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
 