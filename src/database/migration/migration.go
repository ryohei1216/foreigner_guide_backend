package migration

import (
	"foreigner_guide/src/database"
	"foreigner_guide/src/models"
)
func Init() {
	db := database.Db
	//DBリセット時にコメント外す
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

}