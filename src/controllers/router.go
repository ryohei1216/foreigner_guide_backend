package controllers

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	// CORS 対応
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	}))

	router.POST("/signUp", CreateUser)
	router.POST("/signIn", GetSignInUser)
	router.GET("/country", GetCountryImages)
	router.GET("/country_wiki", GetCountryWiki)
	router.GET("/user_id", GetUserById)
	router.GET("/users", GetAllUsers)
	router.GET("/users_area", GetUsersByArea)
	router.GET("/users_apply", GetApplyUsers)
	router.POST("/getMessages", GetMessages)
	router.POST("/uploadImage", UploadImage)
	router.POST("/applyMatch", SaveApplyUser)
	router.POST("/saveMessages", SaveMessages)

	router.Run(":" + os.Getenv("PORT"))
	// router.Run(":8080")
}
