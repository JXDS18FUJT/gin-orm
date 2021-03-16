package router

import (
	"github.com/gin-gonic/gin"
	"orm.com/m/v2/controller"
)

func UploadRouter(Router *gin.RouterGroup) {
	// v1
	uploadGroup := Router.Group("upload")
	{
		// 待办事项
		// 添加
		uploadGroup.POST("/image", controller.UploadImage)
		// 查看所有的待办事项

		// 修改某一个待办事项
		//v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		//v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
}
