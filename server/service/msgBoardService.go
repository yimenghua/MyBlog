package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/dao"
	"server/entity"
	"server/errors"
)

var MsgBoardSever = new(MsgBoardService)

type MsgBoardService struct {
}

func (s *MsgBoardService) MsgBoardInsert(context *gin.Context) {
	var message entity.Message
	err := context.ShouldBindJSON(&message)
	if err != nil {
		log.Println(errors.BadParameter, err)
		context.JSON(errors.HttpBadQuery, gin.H{
			"msg": errors.BadParameter,
		})
	} else {
		err = dao.MsgBoardDAO.MsgBoardInsert(message)
		if err != nil {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.DatabaseOperationFailed,
			})
		} else {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.Success,
			})
		}
	}
}

func (s *MsgBoardService) MsgBoardDelete(context *gin.Context) {
	var message entity.Message
	err := context.ShouldBindJSON(&message)
	if err != nil {
		log.Println(errors.BadParameter, err)
		context.JSON(errors.HttpBadQuery, gin.H{
			"msg": errors.BadParameter,
		})
	} else {
		err = dao.MsgBoardDAO.MsgBoardDelete(message)
		if err != nil {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.DatabaseOperationFailed,
			})
		} else {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.Success,
			})
		}
	}
}

func (s *MsgBoardService) MsgBoardUpdate(context *gin.Context) {
	var message entity.Message
	err := context.ShouldBindJSON(&message)
	if err != nil {
		log.Println(errors.BadParameter, err)
		context.JSON(errors.HttpBadQuery, gin.H{
			"msg": errors.BadParameter,
		})
	} else {
		err = dao.MsgBoardDAO.MsgBoardUpdate(message)
		if err != nil {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.DatabaseOperationFailed,
			})
		} else {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.Success,
			})
		}
	}
}
func (s *MsgBoardService) MsgBoardSelect(context *gin.Context) {
	var message entity.Message
	err := context.ShouldBindJSON(&message)
	if err != nil {
		log.Println(errors.BadParameter, err)
		context.JSON(errors.HttpBadQuery, gin.H{
			"msg": errors.BadParameter,
		})
	} else {
		var msgs []entity.Message
		msgs, err = dao.MsgBoardDAO.MsgBoardSelect(message)
		if err != nil {
			context.JSON(errors.HttpOk, gin.H{
				"msg": errors.DatabaseOperationFailed,
			})
		} else {
			context.JSON(errors.HttpOk, gin.H{
				"msg":  errors.Success,
				"data": msgs,
			})
		}
	}
}
