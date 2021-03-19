package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {

	form, err := c.MultipartForm()
	files := form.File["file"]

	// filesContent := files[0].Content
	// fmt.Sprintln(filesContent)
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var dst string
	for _, f := range files {
		dst = fmt.Sprintf(`E:\gin\gin-orm\upload\` + f.Filename)
		c.SaveUploadedFile(f, dst)

		// srcByte, err := ioutil.ReadFile(dst)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// res := base64.StdEncoding.EncodeToString(srcByte)
		// fmt.Println(res)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})

}
func UploadImagetoBase64(c *gin.Context, index int) (err error, res string, fileName string) {

	form, err := c.MultipartForm()
	files := form.File["file"]

	// filesContent := files[0].Content
	// fmt.Sprintln(filesContent)
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return err, "", ""
	}
	var dst string
	for _, f := range files {
		dst = fmt.Sprintf(`E:\gin\gin-orm\upload\` + f.Filename)
		c.SaveUploadedFile(f, dst)

		srcByte, err := ioutil.ReadFile(dst)
		if err != nil {
			log.Fatal(err)
		}
		res = base64.StdEncoding.EncodeToString(srcByte)
		fileName = f.Filename

	}
	return nil, res, fileName

}
