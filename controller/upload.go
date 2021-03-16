package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {

	form, err := c.MultipartForm()
	files := form.File["file"]
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var dst string
	for _, f := range files {
		dst = fmt.Sprintf(`D:\gin\go-orm\upload\` + f.Filename)
		c.SaveUploadedFile(f, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
