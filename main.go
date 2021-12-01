package main

import "github.com/gin-gonic/gin"

func main() {
	// 返回的Engine 实例里面 用了两个中间件 Logger(输出请求日志) , Recovery(异常捕获)
	r := gin.Default()

	// 定义一个路由
	r.GET("/ping", func(c *gin.Context) {
		// gin.H 是个map类型
		c.JSON(200, gin.H{"message": "pong"})
	})
	// 启动后命令行会提示 Listening and serving HTTP on :8080
	//  默认端口 8080
	r.Run()
}
