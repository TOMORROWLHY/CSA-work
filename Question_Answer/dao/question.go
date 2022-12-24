package dao

import (
	"GO_study/CSA_golong/Question_Answer/model"
	"log"
)

// 插入问题
func InserQuestion(q model.Question) (err error) {
	tx := DB.Create(&q)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	return err
}

// 删除问题
// 通过qid来删除问题
func RemoveQuestion(qid int) (err error) {
	var q model.Question
	tx := DB.Where("qid=?", qid).Delete(&q)
	err = tx.Error
	if err != nil {
		log.Printf("the err %#v", err)
		return
	}
	return err
}

// 修改问题
// 通过qid来修改问题
func ChangeQuestion(qid int, changedquestion string) (err error) {
	var q model.Question
	tx := DB.Model(&q).Where("qid=?", qid).Update("question", changedquestion)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	return err
}

// 通过问题查找问题的qid
func SelectQidByQuestion(question string) (qid int, err error) {
	var q model.Question
	tx := DB.Where("question = ?", question).Find(&q)
	err = tx.Error
	if err != nil {
		log.Printf("the err: %#v", err)
		return
	}
	qid = q.Qid
	return qid, err
}

// 通过uid确定qid是否属于这个用户，即是否是这个用户提出的问题
func SelectQidByUid(uid int) (qids []int, err error) {
	//暂时设定为10个容量上限
	qs := make([]model.Question, 10)
	tx := DB.Where("uid=?", uid).Find(&qs)
	err = tx.Error
	if err != nil {
		log.Printf("the error: %#v", err)
		return
	}
	//将查找出的qid存入qids中
	for i := 0; i < 10; i++ {
		qids = append(qids, qs[i].Qid)
	}
	return qids, err
}

// 查找问题
// 通过uid来查找自己发布的所有问题
func SelectQuestionByUid(uid int) (qs []model.Question, err error) {
	tx := DB.Where("uid=?", uid).Find(&qs)
	err = tx.Error
	if err != nil {
		log.Printf("the err:%#v", err)
		return
	}
	return qs, err
}
