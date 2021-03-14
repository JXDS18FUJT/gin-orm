package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"orm.com/m/v2/model"
)

func CreateClassmate(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var classmate model.Classmate
	c.BindJSON(&classmate)
	// 2. 存入数据库
	err := model.CreateClassmate(&classmate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, classmate)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}

}
func GetAllClassmate(c *gin.Context) {
	// 查询todo这个表里的所有数据
	classmate, err := model.GetAllClassmate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": classmate})
	}

}

func GetClassmate(c *gin.Context) {
	// 查询todo这个表里的所有数据
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	classmate, err := model.GetClassmate(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": classmate})
	}

}

func DeleteClassmate(c *gin.Context) {
	// 查询todo这个表里的所有数据
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := model.DeleteClassmate(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}
func UpdateClassmate(c *gin.Context) {
	// 查询todo这个表里的所有数据
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := model.GetClassmate(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = model.UpdateClassmate(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
