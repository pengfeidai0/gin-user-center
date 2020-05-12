package model

import (
	"gin-user-center/app/database/mysql"
	"time"
)

type User struct {
	UserId    int        `json:"userId" gorm:"primary_key;auto_increment" description:"用户id"`
	Phone     string     `json:"phone" gorm:"size:32;unique_index;not null" description:"手机号"`
	Name      string     `json:"name" gorm:"size:32;not null" description:"用户名"`
	Password  string     `json:"password" gorm:"type:char(32);not null" description:"密码"`
	Salt      string     `json:"sale" gorm:"type:char(16);not null" description:"密码salt"`
	Avatar    string     `json:"avatar" gorm:"size:64" description:"头像"`
	CreatedAt LocalTime  `json:"createdAt" description:"创建时间"`
	UpdatedAt LocalTime  `json:"ureatedAt" description:"更新时间"`
	DeletedAt *time.Time `json:"deletedAt" description:"删除时间"`
}

// 是否存在
func IsExist(phone string) bool {
	var count int
	mysql.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
