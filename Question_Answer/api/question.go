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

// 发布自己的问题
func PutQuestion(c *gin.Context) {
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	question := c.PostForm("question")
	if question == "" {
		tools.RespParamErr(c)
		return
	}
	err = service.CreateQuestion(model.Question{
		Uid:      uid,
		Question: question,
	})
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	qid, _ := service.SelectQid(question)
	qidstring := strconv.Itoa(qid)
	tools.RespOK(c, "发布问题成功"+"请记住您的Qid:"+qidstring)

}

// 修改自己的问题
func ChangeQuestion(c *gin.Context) {
	//判断是否登录
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	qidstring := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidstring)
	newquestion := c.PostForm("new")
	//判断是否是为有效信息
	if qidstring == "" || newquestion == "" {
		tools.RespParamErr(c)
		return
	}
	//判断填入问题是否属于该用户
	flag := service.CheckQIfBelongToYou(uid, qid)
	if !flag {
		tools.NormalErr(c, 403, "您没有修改该问题的权限！")
		return
	}
	err = service.ChangeQuestion(qid, newquestion)
	if err != nil {
		tools.NormalErr(c, 500, "修改问题失败！")
	}
	tools.RespOK(c, "已成功修改该问题！")
}

// 删除自己的问题
func DeleteQuestion(c *gin.Context) {
	//判断是否登录
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	qidstring := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidstring)
	//判断是否为有效信息
	if qidstring == "" {
		tools.RespParamErr(c)
		return
	}
	//判断填入问题是否属于该用户
	flag := service.CheckQIfBelongToYou(uid, qid)
	if !flag {
		tools.NormalErr(c, 403, "您没有删除该问题的权限！")
		return
	}
	err = service.DeleteQuestion(qid)
	if err != nil {
		tools.NormalErr(c, 500, "删除问题失败！")
	}
	tools.RespOK(c, "已成功删除该问题！")
}

// 查找自己发布的问题
func FindAllQuestions(c *gin.Context) {
	//判断是否登录
	uid, err := middleware.Checklogin(c)
	if err != nil {
		//401 => 未认证，没有登录网站
		tools.NormalErr(c, 401, "请先进行登录！")
	}
	qs, err := service.FindQuestion(uid)
	if err != nil {
		tools.NormalErr(c, 500, "未能找到该问题！")
		return
	}
	tools.RespOK(c, "已为您找到您所提出的所有问题")
	//将用户的问题输出
	for i := 0; i < 10; i++ {
		c.JSON(http.StatusOK, gin.H{
			"Qid":      qs[i].Qid,
			"Question": qs[i].Question,
		})
	}
}
