package router

import (
	"github.com/gin-gonic/gin"
	"server/service"
)

func initMsgBoardRouter(Router *gin.Engine) {
	msgBoardRouter := Router.Group("msgboard")
	msgBoardRouter.POST("insert", service.MsgBoardSever.MsgBoardInsert)
	msgBoardRouter.POST("delete", service.MsgBoardSever.MsgBoardDelete)
	msgBoardRouter.POST("update", service.MsgBoardSever.MsgBoardUpdate)
	msgBoardRouter.POST("select", service.MsgBoardSever.MsgBoardSelect)
}
