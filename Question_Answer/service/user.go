package service

import (
	"GO_study/CSA_golong/Question_Answer/dao"
	"GO_study/CSA_golong/Question_Answer/model"
)

// 创建用户
func CreatUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

// 通过用户查找用户
func SearchUserByUsername(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(name)
	return
}

// 修改密码
func Change(username, newpass string) error {
	err := dao.ChangePass(username, newpass)
	return err
}
