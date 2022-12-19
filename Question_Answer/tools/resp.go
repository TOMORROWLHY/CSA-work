package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一信息反馈模板
type respTemplate struct {
	Status int    `json:"Status"`
	info   string `json:"Info"`
}

var OK = respTemplate{
	Status: 200,
	info:   "success",
}

var ParamError = respTemplate{
	Status: 300,
	info:   "params error",
}

var InternalError = respTemplate{
	Status: 500,
	info:   "internal error",
}

func RespOK(c *gin.Context, info string) { //正确
	OK.info = info
	c.JSON(http.StatusOK, OK)
}
func RespParamErr(c *gin.Context) { //错误
	c.JSON(http.StatusBadRequest, ParamError)
}

func RsepInternalErr(c *gin.Context) { //连接错误
	c.JSON(http.StatusInternalServerError, InternalError)
}

func NormalErr(c *gin.Context, status int, info string) { //其他错误
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
