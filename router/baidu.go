package router

import (
	"github.com/gin-gonic/gin"
	"orm.com/m/v2/controller"
)

func BaiduRouter(Router *gin.RouterGroup) {
	// v1
	baiduRouterGroup := Router.Group("baidu")
	{
		// 待办事项
		// 添加
		//classmateRouterGroup.POST("/todo", controller.CreateV1Todo)
		// 查看所有的待办事项
		baiduRouterGroup.POST("/BaiduWebimage", controller.BaiduWebimage)

		// 修改某一个待办事项
		//v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		//v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
}
