package api

import (
	"GO_study/CSA_golong/Question_Answer/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//解决跨域问题
	r.Use(middleware.CORS())
	//用户路由组
	u := r.Group("/user")
	u.POST("/register", Register)
	u.POST("login", Login)
	u.PUT("/change", ChangePass)
	//问题路由组
	q := r.Group("/question")
	q.POST("/post", PutQuestion)
	q.PUT("/change", ChangeQuestion)
	q.DELETE("/delete", DeleteQuestion)
	q.GET("/find", FindAllQuestions)
	//回答路由组
	a := r.Group("/answer")
	a.POST("/post", PostAnswer)
	a.PUT("/change", ChangeAnswer)
	a.DELETE("/delete", DeleteAnswer)
	a.GET("find", FindAllAnswers)

	r.Run(":8080")

}
