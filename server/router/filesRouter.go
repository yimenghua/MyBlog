package router

import (
	"github.com/gin-gonic/gin"
)

func initFilesRouter(Router *gin.Engine) {
	filesRouter := Router.Group("files")
	filesRouter.Static("front", "../files/front")
	filesRouter.Static("upload", "../files/upload")
}
