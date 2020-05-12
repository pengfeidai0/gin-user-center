package controller

import (
	"gin-user-center/app/common"
	"gin-user-center/app/middleware"
	"gin-user-center/app/model"
	"gin-user-center/app/schema"
	"gin-user-center/app/service"

	"github.com/gin-gonic/gin"
)

var logger = common.Logger

/**
* 新增用户
 */
func AddUser(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

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
	u, err := service.AddUser(user)
	if err != nil {
		logger.Error("controller addUser error:", err)
		ctx.Response(common.ERROR, err.Error(), nil)
		return
	}

	data := map[string]interface{}{
		"userId": u.UserId,
		"name":   u.Name,
		"avatar": u.Avatar,
	}
	ctx.Response(common.SUCCESS, nil, data)
}
