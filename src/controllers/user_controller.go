package controllers

import (
	"foreigner_guide/src/models"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RandomString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}

func CreateUser(c *gin.Context) {
	var user models.User
	user.Id = RandomString(10)
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

func GetSignInUser(c *gin.Context){
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
	}
	getUser, err := user.Get()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"getUser": getUser,
	})
}

func GetUsersByArea (c *gin.Context) {
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