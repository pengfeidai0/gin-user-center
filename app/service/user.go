package service

import (
	"errors"
	"gin-user-center/app/common"
	"gin-user-center/app/model"
)

var logger = common.Logger

func AddUser(u model.User) error {
	logger.Info("service AddUser param:", u)
	// var user model.User
	// 手机号是否已注册
	isExist := model.IsExist(u.Phone)
	if isExist {
		return errors.New(common.PHONE_EXIST)
	}
	return nil
}
