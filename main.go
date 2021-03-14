package main

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
	"orm.com/m/v2/dao"
	"orm.com/m/v2/initialize"
	"orm.com/m/v2/model"
	"orm.com/m/v2/setting"
)

func main() {
	//加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	r := initialize.InitRouter()
	logger, _ := zap.NewProduction()
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//添加一个ginzap中间件，它：
	//记录所有请求，如组合的访问和错误日志。
	//记录到stdout。
	//RFC3339，UTC时间格式。
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	//将所有死机记录到错误日志中
	//stack表示是否输出堆栈信息。
	r.Use(ginzap.RecoveryWithZap(logger, true))
	// 模型绑定
	dao.DB.AutoMigrate(&model.Classmate{})
	defer dao.Close() // 程序退出关闭数据库连接
	//静态资源处理
	// r.Static("/static", "./static")
	//静态页面处理
	// r.LoadHTMLGlob("templates/*")
	//根路径访问登录页面
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8001")

}
