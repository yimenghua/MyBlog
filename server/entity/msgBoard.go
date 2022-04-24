package entity

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	UserName string `json:"UserName"`
	Msg      string `json:"Msg"`
}

func (m Message) TableName() string {
	return "msgboard"
}
