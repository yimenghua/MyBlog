package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"server/config"
	"server/errors"
)

var MySqlDB *gorm.DB

func InitDb() {
	MySqlDB = getMysqlDb()
}

func getMysqlDb() *gorm.DB {
	mysqlConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUserName, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDbName)
	db, err := gorm.Open("mysql", mysqlConn)
	if err != nil {
		log.Fatalln(errors.MysqlConnectFail, err)
		return nil
	} else {
		log.Println(errors.MysqlConnectSuccess)
		return db
	}
}
