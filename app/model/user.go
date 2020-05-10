package model

import (
	"gin-user-center/app/database/mysql"
	"time"
)

// const (
// 	db         = "gin-user-center"
// 	collection = "bas_user"
// )

// var db = mysql.DB

type User struct {
	UserId   int       `json:"id" gorm:"primary_key;auto_increment" description:"用户id"`
	Phone    string    `json:"phone" gorm:"size:32;unique_index;not null" description:"手机号"`
	Name     string    `json:"name" gorm:"size:32;not null" description:"用户名"`
	Password string    `json:"password" gorm:"type:char(32);not null" description:"密码"`
	Salt     string    `json:"sale" gorm:"type:char(16);not null" description:"密码salt"`
	Avatar   string    `json:"avatar" gorm:"size:64" description:"头像"`
	CreateAt time.Time `json:"createAt" gorm:"type:datetime" description:"创建时间"`
	UpdateAt time.Time `json:"updateAt" gorm:"type:datetime" description:"更新时间"`
	Invaild  string    `json:"invaild" gorm:"type:char(1);default:'N'" description:"更新时间"`
}

// 是否存在
func IsExist(phone string) bool {
	var count int
	mysql.DB.Model(&User{}).Count(&count)
	return count > 0
}
