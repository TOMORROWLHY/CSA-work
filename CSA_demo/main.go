package main

import (
	"CSA_demo/api"
	"CSA_demo/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
