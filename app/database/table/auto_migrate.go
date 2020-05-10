package table

import (
	"gin-user-center/app/common"
	"gin-user-center/app/database/mysql"
	"gin-user-center/app/model"
)

var logger = common.Logger

func Init() {
	db := mysql.DB
	db.AutoMigrate(&model.User{})
	logger.Info("AutoMigrate tables success.")
}
