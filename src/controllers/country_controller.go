package controllers

import (
	"fmt"
	"foreigner_guide/src/controllers/bing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCountryImages(c *gin.Context) {
	q := c.Query("q")
	ans := bing.GetImages(q)
	fmt.Println(ans)
	c.JSON(http.StatusOK, gin.H{
		"countries": ans,
	})
}