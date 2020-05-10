package controller

import (
	"gin-user-center/app/common"
	"gin-user-center/app/model"
	"gin-user-center/app/schema"
	"gin-user-center/app/service"
	"gin-user-center/app/util"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	ctx := util.Context{Ctx: c}

	var p schema.AddUser
	if err := ctx.ValidateJSON(&p); err != nil {
		return
	}
	user := model.User{
		Phone:    p.Phone,
		Name:     p.Name,
		Password: p.Password,
		Avatar:   p.Avatar,
	}
	err := service.AddUser(user)
	if err != nil {
		ctx.Response(common.HAVE_EXIST, err.Error(), nil)
	}
}
