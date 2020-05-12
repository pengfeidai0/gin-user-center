package service

import (
	"errors"
	"gin-user-center/app/common"
	"gin-user-center/app/database/mysql"
	"gin-user-center/app/model"
	"gin-user-center/app/util"
)

var logger = common.Logger

/**
* 新增用户
 */
func AddUser(u model.User) (user model.User, err error) {
	// 手机号格式校验
	if false == util.CheckPhone(u.Phone) {
		return user, errors.New(common.INVALID_PHONE)
	}

	// 手机号是否已注册
	isExist := model.IsExist(u.Phone)
	if isExist {
		return user, errors.New(common.PHONE_EXIST)
	}

	u.Salt = util.RandString(16)
	u.Password = util.Md5(u.Password, u.Salt)

	logger.Info("service AddUser param:", u)

	err = mysql.DB.Create(&u).Error
	return u, err
}
