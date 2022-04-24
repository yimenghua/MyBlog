package dao

import (
	"log"
	"server/entity"
)

var MsgBoardDAO = new(MsgBoardDao)

type MsgBoardDao struct {
}

func (d *MsgBoardDao) MsgBoardInsert(msg entity.Message) error {
	err := MySqlDB.Create(&msg).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *MsgBoardDao) MsgBoardDelete(msg entity.Message) error {
	log.Println(msg)
	return nil
}

func (d *MsgBoardDao) MsgBoardUpdate(msg entity.Message) error {
	log.Println(msg)
	return nil
}

func (d *MsgBoardDao) MsgBoardSelect(msg entity.Message) ([]entity.Message, error) {
	var msgs []entity.Message
	err := MySqlDB.Order("created_at asc").Find(&msgs).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return msgs, nil
}
