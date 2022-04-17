package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/files", "../files")

	r.Static("/MyBlog", "../web/dist")

	r.Run(":8080")
}
