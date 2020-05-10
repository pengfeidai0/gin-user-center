package mysql

import (
	"gin-user-center/app/common"
	"gin-user-center/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var logger = common.Logger

func Init() {
	var err error
	config := config.Conf.Mysql
	url := config.User + ":" + config.Password + "@(" + config.Path + ")/" + config.Database + "?" + config.Config
	db, err := gorm.Open(config.Driver, url)
	if err != nil {
		logger.Error("mysql connection error: ", err)
		panic(err)
	}
	// db.DB().SetMaxIdleConns(config.MaxIdleConns)
	// db.DB().SetMaxOpenConns(config.MaxOpenConns)
	// 全局禁用表名复数
	db.SingularTable(true)
	db.LogMode(config.Log)

	logger.Info("mysql connection open to: ", url)
	DB = db
}
