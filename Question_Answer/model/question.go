package model

type Question struct {
	Qid      int    `json:"qid"`
	Uid      int    `json:"uid"`
	Question string `json:"question"`
}
