package api

import (
	"CSA_demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
func register(c *gin.Context) {
	//传入用户名和密码
	name := c.PostForm("username")
	password := c.PostForm("password")
	//判断用户名是否重复
	flag := dao.SelectUser(name)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "the username has already existed",
		})
		return
	}
	//不重复,进行创建
	if name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "name or password is nil ",
		})
		return
	}
	dao.AddUser(name, password)
	c.JSON(http.StatusOK, gin.H{
		"message": "add uesr successfully!",
	})
}

// 登录
func login(c *gin.Context) {
	//传入用户名和密码
	name := c.PostForm("username")
	password := c.PostForm("password")
	//判断用户名是否重复
	flag1 := dao.SelectUser(name)
	if !flag1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user doesn't exists",
		})
		return
	}
	//比较输入密码与用户名是否匹配
	flag2 := dao.SelectPassword(name, password)
	if !flag2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "the password is wrong! please try again!",
		})
		return
	}
	//正确则登录成功  设置cookie
	c.SetCookie("GeekBee_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Log in successfully!",
	})
}
