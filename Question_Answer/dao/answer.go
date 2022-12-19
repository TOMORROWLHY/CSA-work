package dao

import (
	"GO_study/CSA_golong/Question_Answer/model"
	"log"
)

// 添加回答
func InsertAnswer(a model.Answer) (err error) {
	tx := DB.Create(&a)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
	}
	return err
}

// 删除回答
// 通过aid删除问题
func RemoveAnswer(aid int) (err error) {
	var a model.Answer
	tx := DB.Where("aid=?", aid).Delete(&a)
	err = tx.Error
	if err != nil {
		log.Printf("the err %#v", err)
		return
	}
	return err
}

// 修改回答
// 通过aid来修改回答
func ChangeAnswer(aid int, changedanswer string) (err error) {
	var a model.Answer
	tx := DB.Model(&a).Where("aid=?", aid).Update("answer", changedanswer)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	return err
}

// 通过回答查找问题的aid
func SelectAidByAnswer(answer string) (aid int, err error) {
	var a model.Answer
	tx := DB.Select("aid").Where("answer= ?", answer).Find(&a)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	aid = a.Aid
	return aid, err
}

// 通过aid查找该回答的用户
func SelectUerByAid(aid int) (a model.Answer, err error) {
	tx := DB.Select("uid").Where("aid=?", aid).Find(&a)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	return a, err
}

// 通过UID来查找自己的所有回答
func SelectAnswersByUid(uid int) (as []model.Answer, err error) {
	tx := DB.Where("uid=?", uid).Find(&as)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	return as, err
}
