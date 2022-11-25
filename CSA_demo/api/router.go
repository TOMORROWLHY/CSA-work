package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.POST("/register", register) // 注册
	r.POST("/login", login)       // 登录

	r.Run(":8088") // 跑在 8088 端口上
}
