package main

import (
	"GO_study/CSA_golong/Question_Answer/api"
	"GO_study/CSA_golong/Question_Answer/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
