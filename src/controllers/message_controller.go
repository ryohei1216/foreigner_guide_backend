package controllers

import (
	"fmt"
	"foreigner_guide/src/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveMessages(c *gin.Context) {
	var postValues struct {
		SignInInfoId string
		UserId   		 string
		Message 		 string

	}
	c.BindJSON(&postValues)
	var message models.Message
	message.UserId = postValues.UserId
	message.Message = postValues.Message
	message.CreatedAt = time.Now()
	fmt.Println(postValues)

	err := message.CreateBySignInfoId(postValues.SignInInfoId)
	if err != nil {
		log.Println(err)
	}
}

func GetMessages (c *gin.Context) {
	var postValues struct {
		UserId string
		ChatUserId string
	}
	c.BindJSON(&postValues)

	messages, err := models.GetMessagesByIds(postValues.UserId, postValues.ChatUserId)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}