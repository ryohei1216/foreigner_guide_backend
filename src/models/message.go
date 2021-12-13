package models

import (
	"foreigner_guide/src/database"
	"time"
)

type Message struct {
	UserId    string
	Message   string
	CreatedAt time.Time
}

func (m *Message) CreateBySignInfoId(signInInfoId string) (err error){
	db := database.Db
	err = db.Table("messages_"+signInInfoId).Create(m).Error
	return err
}

func GetMessagesByIds(userId string, chatUserId string) (messages []Message, err error) {
	db := database.Db
	messages = []Message{}
	err = db.Table("messages_"+userId).Where("user_id = ?", chatUserId).Find(&messages).Error
	return messages, err
}