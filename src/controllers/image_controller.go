package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// //画像をアップロードするPOSTメソッド
// func uploading(w http.ResponseWriter, r *http.Request){
// 	file, header, err := r.FormFile("uploading")
// 	if err != nil {
// 		fmt.Println("読み込み失敗")
// 	}
// 	defer file.Close()

// 	data, _ := ioutil.ReadAll(file)
// 	newFile, err := os.Create(filepath.Join("./app/views/images/", header.Filename))
// 	if err != nil {
// 		fmt.Println("ファイル作成失敗")
// 		fmt.Println(err)
// 	}
// 	defer newFile.Close()

// 	fileStr := string(data)
// 	_, err = newFile.Write([]byte(fileStr))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func UploadImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	fileHeader := form.File["image"][0]
  file, err := fileHeader.Open()
  if err != nil {
		log.Println(err)
	}
	data, _ := ioutil.ReadAll(file)

  newFile, err := os.Create(filepath.Join("./images",fileHeader.Filename))
  if err != nil {
		fmt.Println("ファイル作成失敗")
		log.Println(err)
	}
	defer newFile.Close()

  fileStr := string(data)
	_, err = newFile.Write([]byte(fileStr))
  if err != nil {
		log.Println(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"result": fileStr})
  // c.File("./images/download1.jpg")
}

func GetImage(imgName string) {

}
