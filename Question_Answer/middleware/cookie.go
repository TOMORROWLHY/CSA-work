package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 检查登录状态并返回uid
func Checklogin(c *gin.Context) (uid int, err error) {
	uidstring, err := c.Cookie("user")
	uid, _ = strconv.Atoi(uidstring)
	if err != nil {
		log.Printf("the error:%#v", err)
		return
	}
	return uid, err
}
