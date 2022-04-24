package main

import (
	_ "github.com/go-sql-driver/mysql"
	"server/dao"
	"server/logs"
	"server/router"
)

func main() {

	logs.InitLog()      //初始日志输出
	dao.InitDb()        //初始化数据库
	router.InitRouter() //初始化路由

	//程序结束前关闭数据库连接
	if dao.MySqlDB != nil {
		defer dao.MySqlDB.Close()
	}

	//启动服务
	if router.Router != nil {
		router.Router.Run(":8080")
	}
}
