package main

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"orm.com/m/v2/dao"
	"orm.com/m/v2/initialize"
	"orm.com/m/v2/middlewares"
	"orm.com/m/v2/model"
	"orm.com/m/v2/setting"
)

func main() {

	//加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	r := gin.New()
	// 开启loger
	logger, _ := zap.NewProduction()
	//添加一个ginzap中间件，它：
	//记录所有请求，如组合的访问和错误日志。
	//记录到stdout。
	//RFC3339，UTC时间格式。
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	//将所有死机记录到错误日志中
	//stack表示是否输出堆栈信息。
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middlewares.Cors())
	initialize.InitRouter(r)

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	//开启中间件 允许使用跨域请求

	// 模型绑定
	dao.DB.AutoMigrate(&model.Classmate{})
	defer dao.Close() // 程序退出关闭数据库连接
	//静态资源处理
	// r.Static("/static", "./static")
	//静态页面处理
	r.LoadHTMLGlob("templates/index.html")
	//根路径访问登录页面
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
