package router

import (
	"github.com/gin-gonic/gin"
	"server/middle"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.Use(middle.Cors())
	Router.Static("/myblog", "../web/dist") //设置我的博客访问路径
	initFilesRouter(Router)                 //设置文件访问路径
	initMsgBoardRouter(Router)              //设置留言板接口路径
}
