package api

import (
	"GO_study/CSA_golong/Question_Answer/model"
	"GO_study/CSA_golong/Question_Answer/service"
	"GO_study/CSA_golong/Question_Answer/tools"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	//判断用户名或者密码是否为空，减少不必要的资源占用
	if username == "" || password == "" {
		tools.RespParamErr(c)
		return
	}
	//判断用户名是否重复
	flag := service.CheckExist(username)
	if flag {
		tools.NormalErr(c, 500, "用户名已存在")
		return
	}
	err := service.CreatUser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespOK(c, "注册成功！")
}

// 用户登录
func Login(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	if username == "" || password == "" {
		tools.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err != sql.ErrNoRows {
			tools.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RsepInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		tools.NormalErr(c, 500, "密码错误！")
		return
	}
	//通过用户名查找出uid
	u, err = service.SearchUserByUsername(username)
	if err != nil {
		if err != sql.ErrNoRows {
			tools.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RsepInternalErr(c)
			return
		}
		return
	}
	tools.RespOK(c, "登录成功")
	uidstring := strconv.Itoa(u.ID)
	c.SetCookie("user", uidstring, 3600, "/", "", false, true)
	//_, err = c.Cookie("user")
	//if err != nil {
	//	log.Printf("set cookie err: %#v", err)
	//	return
	//}
}

// 修改密码
func ChangePass(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	newpass := c.PostForm("newpassword")
	if username == "" || password == "" {
		tools.RespParamErr(c)
		return
	}
	_, err := service.SearchUserByUsername(username)
	if err != nil {
		if err != sql.ErrNoRows {
			tools.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RsepInternalErr(c)
			return
		}
		return
	}
	err = service.Change(username, newpass)
	if err != nil {
		tools.NormalErr(c, 300, "修改密码失败！")
		return
	}
	tools.RespOK(c, "修改密码成功！")
}
