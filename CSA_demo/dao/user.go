package dao

import (
	"CSA_demo/model"
	"fmt"
)

// 添加用户
func AddUser(name, password string) {

	u := model.User{
		Name:     name,
		Password: password,
	}
	DB.Create(&u)
}

// 查找用户名
func SelectUser(name string) bool {

	var u model.User
	a := DB.Select("name").Where("name = ?", name).Find(&u)
	if a.Error != nil {
		fmt.Printf("%#v", a.Error)
		panic(a.Error)
	}

	fmt.Println(u.Name)
	if u.Name != "" {
		return true
	}
	return false
}

// 从数据表中查找密码
func SelectPassword(name, password string) bool {

	var u model.User
	DB.Where("password = ? AND name = ?", password, name).Find(&u)
	if u.Password != password {
		return false
	}
	return true
}
