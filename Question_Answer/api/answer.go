package api

import (
	"GO_study/CSA_golong/Question_Answer/middleware"
	"GO_study/CSA_golong/Question_Answer/model"
	"GO_study/CSA_golong/Question_Answer/service"
	"GO_study/CSA_golong/Question_Answer/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 发布回答
func PostAnswer(c *gin.Context) {
	//判断是否登录
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	qidString := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidString)
	answer := c.PostForm("answer")
	if qidString == "" || answer == "" {
		tools.RespParamErr(c)
		return
	}
	err = service.CreateAnswer(model.Answer{
		Uid:    uid,
		Qid:    qid,
		Answer: answer,
	})
	if err != nil {
		tools.NormalErr(c, 500, "发表回答失败！")
		return
	}
	aid, _ := service.SelectAid(answer)
	aidstring := strconv.Itoa(aid)
	tools.RespOK(c, "发表回答成功"+" 请记住您的Aid:"+aidstring)
}

// 删除回答
func DeleteAnswer(c *gin.Context) {
	//判断是否登录
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	aidString := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidString)
	//判断是否为有效信息
	if aidString == "" {
		tools.RespParamErr(c)
		return
	}
	//判断是否是该用户提出的回答
	flag := service.CheckAIfBelongToUser(uid, aid)
	if !flag {
		tools.NormalErr(c, 403, "您没有删除该回答的权限！")
		return
	}
	err = service.DeleteAnswer(aid)
	if err != nil {
		tools.NormalErr(c, 500, "删除回答失败！")
	}
	tools.RespOK(c, "已成功删除该回答！")
}

// 修改回答
func ChangeAnswer(c *gin.Context) {
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	changedanswer := c.PostForm("changedanswer")
	aidString := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidString)
	//判断是否为有效信息
	if aidString == "" || changedanswer == "" {
		tools.RespParamErr(c)
		return
	}
	//判断是否是该用户提出的回答
	flag := service.CheckAIfBelongToUser(uid, aid)
	if !flag {
		tools.NormalErr(c, 403, "您没有删除该回答的权限！")
		return
	}
	err = service.ChangeAnswer(aid, changedanswer)
	if err != nil {
		tools.NormalErr(c, 500, "修改回答失败！")
	}
	tools.RespOK(c, "已成功修改该回答！")
}

// 查找自己的所有回答
func FindAllAnswers(c *gin.Context) {
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	as, err := service.FindAnswer(uid)
	if err != nil {
		tools.NormalErr(c, 500, "未能找到该回答")
		return
	}
	tools.RespOK(c, "已为您找到您的所有回答！")
	//将用户的所有回答输出
	for i := 0; i < 10; i++ {
		c.JSON(http.StatusOK, gin.H{
			"Qid":    as[i].Qid,
			"Aid":    as[i].Aid,
			"Answer": as[i].Answer,
		})
	}
}
