package router

import (
	"github.com/gin-gonic/gin"
	"orm.com/m/v2/controller"
)

func ClassmateRouter(Router *gin.RouterGroup) {
	// v1
	classmateRouterGroup := Router.Group("classmate")
	{
		// 待办事项
		// 添加
		//classmateRouterGroup.POST("/todo", controller.CreateV1Todo)
		// 查看所有的待办事项
		classmateRouterGroup.POST("/todo", controller.CreateClassmate)
		classmateRouterGroup.GET("/todos", controller.GetAllClassmate)
		classmateRouterGroup.GET("/todo", controller.GetClassmate)
		classmateRouterGroup.DELETE("/todo/:id", controller.DeleteClassmate)
		classmateRouterGroup.PUT("/todo/:id", controller.UpdateClassmate)
		// 修改某一个待办事项
		//v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		//v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
}
