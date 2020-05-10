package service

import (
	"errors"
	"gin-user-center/app/common"
	"gin-user-center/app/database/mysql"
	"gin-user-center/app/model"
)

var logger = common.Logger

func AddUser(u model.User) (err error, user model.User) {
	logger.Info("service AddUser param:", u)
	// 手机号是否已注册
	isExist := model.IsExist(u.Phone)
	if isExist {
		return errors.New(common.PHONE_EXIST), user
	}
	err = mysql.DB.Create(&u).Error
	return err, u
}
