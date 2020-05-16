package service

import (
	"errors"
	"gin-user-center/app/common"
	"gin-user-center/app/database/mysql"
	"gin-user-center/app/model"
	"gin-user-center/app/util"
)

/**
* 用户登录
 */
func Login(phone string, password string) (user model.User, err error) {
	logger.Infof("service login phone: %s, password: %s", phone, password)
	// 手机号格式校验
	if false == util.CheckPhone(phone) {
		return user, errors.New(common.INVALID_PHONE)
	}

	err = mysql.DB.Where(&model.User{Phone: phone}).First(&user).Error
	if err != nil {
		logger.Error("service login error:", err)
		return user, errors.New(common.PHONE_NOT_EXIST)
	}

	if util.Md5(password, user.Salt) != user.Password {
		return user, errors.New(common.PASSWORS_ERROR)
	}

	return user, err
}
