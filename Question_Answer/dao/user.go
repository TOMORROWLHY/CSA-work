package dao

import (
	"GO_study/CSA_golong/Question_Answer/model"
	"fmt"
	"log"
)

// 插入用户
func InsertUser(u model.User) (err error) {
	tx := DB.Create(&u)
	err = tx.Error
	log.Printf("the err: %#v", err)
	return err
}

// 通过用户名查找用户
func SearchUserByUsername(name string) (u model.User, err error) {
	tx := DB.Select("name").Where("name=?", name).Find(&u)
	err = tx.Error
	if err != nil {
		log.Printf("the error: %#v", err)
		panic(err)
	}
	fmt.Println(u.Username)
	return u, err
}

// 修改用户密码
func ChangePass(username string, newpass string) (err error) {
	var u model.User
	tx := DB.Model(&u).Where("username=?", username).Update("password", newpass)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", tx.Error)
		return
	}
	return err
}
