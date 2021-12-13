package controllers

import (
	"fmt"
	"foreigner_guide/src/controllers/api/bing"
	"foreigner_guide/src/controllers/api/wiki"
	"net/http"

	"github.com/gin-gonic/gin"
)

//countryのimage情報取得
func GetCountryImages(c *gin.Context) {
	q := c.Query("q")
	ans := bing.GetImages(q)
	fmt.Println(ans)
	c.JSON(http.StatusOK, gin.H{
		"countries": ans,
	})
}

//countryのwiki情報取得
func GetCountryWiki(c *gin.Context) {
	q := c.Query("q")
	ans := wiki.GetWiki(q)
	fmt.Println(ans)
	c.JSON(http.StatusOK, gin.H{
		"wiki": ans,
	})
}
