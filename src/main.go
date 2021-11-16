package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":" + os.Getenv("PORT"))
}