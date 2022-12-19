package service

import (
	"GO_study/CSA_golong/Question_Answer/dao"
	"GO_study/CSA_golong/Question_Answer/model"
)

// 创建自己的问题
func CreateQuestion(q model.Question) (err error) {
	err = dao.InserQuestion(q)
	return err
}

// 修改自己的问题
func ChangeQuestion(qid int, newquestion string) (err error) {
	err = dao.ChangeQuestion(qid, newquestion)
	return err
}

// 删除自己的问题
func DeleteQuestion(qid int) (err error) {
	err = dao.RemoveQuestion(qid)
	return err
}

// 查找自己发布的问题
func FindQuestion(uid int) (qs []model.Question, err error) {
	qs, err = dao.SelectQuestionByUid(uid)
	return qs, err
}

// 查找问题的qid
func SelectQid(question string) (qid int, err error) {
	qid, err = dao.SelectQidByQuestion(question)
	return qid, err
}

// 判断输入的qid是否是该用户的uid
func CheckQIfBelongToYou(uid, qid int) (flag bool) {
	qids, _ := dao.SelectQidByUid(uid)
	for i := 0; i < 10; i++ {
		if qids[i] == qid {
			flag = true
			break
		} else {
			flag = false
		}
	}
	return flag
}
