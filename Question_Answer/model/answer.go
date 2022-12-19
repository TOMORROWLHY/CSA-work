package model

type Answer struct {
	Aid    int    `json:"aid"`
	Uid    int    `json:"uid"`
	Qid    int    `json:"qid"`
	Answer string `json:"answer"`
}
