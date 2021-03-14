package router

import (
	"github.com/gin-gonic/gin"
	"orm.com/m/v2/controller"
)

func V1Router(Router *gin.RouterGroup) {
	// v1
	v1Group := Router.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateV1Todo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.CreateV1Todo)
		// 修改某一个待办事项
		//v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		//v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
}
