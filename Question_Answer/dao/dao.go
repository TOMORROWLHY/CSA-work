package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@lhy20031018.@tcp(127.0.0.1:3306)/csaqa?charset=utf8mb4&&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err: %#v", err)
	}
	DB = db
}
