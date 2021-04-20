package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type Word struct {
	Words string
}
type webimageRes struct {
	Log_id           int    `json:"log_id"`
	Direction        int    `json:"direction"`
	Words_result_num int    `json:"words_result_num"`
	Words_result     []Word `json:"words_result"`
}
func BaiduWebimage(c *gin.Context) {
	var __dirname = "E:\\gin\\gin-orm\\txt\\"
	// var accessToken = "24.2f5612a038f5c41bf4b370fceca502fc.2592000.1618729517.282335-23829922"
	_, uploadImagetoBase64, fileName := UploadImagetoBase64(c, 0)
	data := make(url.Values)
	data["image"] = []string{uploadImagetoBase64}
	data["detect_direction"] = []string{"true"}
	//把post表单发送给目标服务器
	resp, err := http.PostForm("https://aip.baidubce.com/rest/2.0/ocr/v1/webimage?access_token=24.63b996af3629d9380f346fb432d7ba79.2592000.1621492446.282335-23829922", data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "err",
		})
		return
	}
	//关闭请求
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var webimageResult webimageRes
	jsonErr := json.Unmarshal([]byte(string(body)), &webimageResult)
	if jsonErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "json失败了",
		})
		return

	}
	var txt string
	for _, txtLine := range webimageResult.Words_result {
		txt = txt + txtLine.Words + "\r\n"
	}
	fmt.Println(txt)
	//创建txt
	f, err := os.Create(__dirname + fileName + ".txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(txt))

	}
	c.JSON(http.StatusOK, gin.H{
		"message": webimageResult,
	})

}
