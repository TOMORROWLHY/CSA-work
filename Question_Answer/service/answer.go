package service

import (
	"GO_study/CSA_golong/Question_Answer/dao"
	"GO_study/CSA_golong/Question_Answer/model"
	"log"
)

// 增
func CreateAnswer(a model.Answer) (err error) {
	err = dao.InsertAnswer(a)
	return err
}

// 删
func DeleteAnswer(aid int) (err error) {
	err = dao.RemoveAnswer(aid)
	return err
}

// 查
func FindAnswer(uid int) (as []model.Answer, err error) {
	as, err = dao.SelectAnswersByUid(uid)
	return as, err
}

// 改
func ChangeAnswer(aid int, changedanswer string) (err error) {
	err = dao.ChangeAnswer(aid, changedanswer)
	return err
}

// 查找问题的aid
func SelectAid(answer string) (aid int, err error) {
	aid, err = dao.SelectAidByAnswer(answer)
	return aid, err
}

// 判断问题是否属于该用户
func CheckAIfBelongToUser(uid, aid int) (flag bool) {
	a, err := dao.SelectUerByAid(aid)
	if err != nil {
		log.Printf("the err: %#v", err)
	}
	if uid != a.Uid {
		flag = false
	} else {
		flag = true
	}
	return flag
}
