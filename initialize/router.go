package initialize

import (
	"orm.com/m/v2/controller"
	"orm.com/m/v2/router"
	"orm.com/m/v2/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Base)
	PublicGroup := r.Group("")
	{
		router.V1Router(PublicGroup)        // v1基础路由
		router.ClassmateRouter(PublicGroup) // v1基础路由
		router.UploadRouter(PublicGroup)
		router.BaiduRouter(PublicGroup)

	}

	// PrivateGroup := r.Group("")

	return r

}
