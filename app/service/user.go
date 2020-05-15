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
	logger.Info("service AddUser param:", u)
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

	err = mysql.DB.Create(&u).Error
	return u, err
}

/*
 * 修改密码
 */
func UpdatePassword(phone, oldPassword, newPassword string) (err error) {
	logger.Infof("service AddUser UpdatePassword: %s %s %s", phone, oldPassword, newPassword)
	var user model.User
	// 手机号格式校验
	if false == util.CheckPhone(phone) {
		return errors.New(common.INVALID_PHONE)
	}

	// 查找用户信息
	err = mysql.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return errors.New(common.PHONE_NOT_EXIST)
	}

	if util.Md5(oldPassword, user.Salt) != user.Password {
		return errors.New(common.OLD_PASSWORS_ERROR)
	}

	if oldPassword == newPassword {
		return errors.New(common.NEW_PASSWORD_SAME_AS_OLD)
	}

	password := util.Md5(newPassword, user.Salt)

	// 更新密码
	err = mysql.DB.Model(&user).Where("phone = ?", phone).Update("password", password).Error
	if err != nil {
		logger.Error("service UpdatePassword error:", err)
		return errors.New(common.UPDATE_PASSWORD_FAILD)
	}
	return nil
}

/**
 * 修改头像
 */
// func UploadAvatar(fileName string) error {
// 	var user model.User
// 	// 更新密码
// 	err = mysql.DB.Model(&user).Where("phone = ?", phone).Update("password", password).Error
// 	if err != nil {
// 		logger.Error("service UploadAvatar error:", err)
// 		return errors.New(common.UPDATE_PASSWORD_FAILD)
// 	}
// }
