package model

type User struct {
	Name     string `gorm:"primary_key"`
	Password string
}
