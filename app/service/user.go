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
func AddUser(u model.User) (err error, user model.User) {

	// 手机号格式校验
	if false == util.CheckPhone(u.Phone) {
		return errors.New(common.INVALID_PHONE), user
	}

	// 手机号是否已注册
	isExist := model.IsExist(u.Phone)
	if isExist {
		return errors.New(common.PHONE_EXIST), user
	}

	u.Salt = util.RandString(16)
	u.Password = util.Md5(u.Password, u.Salt)

	logger.Info("service AddUser param:", u)

	err = mysql.DB.Create(&u).Error
	return err, u
}
