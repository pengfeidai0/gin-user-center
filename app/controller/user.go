package controller

import (
	"gin-user-center/app/common"
	"gin-user-center/app/middleware"
	"gin-user-center/app/schema"
	"gin-user-center/app/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var logger = common.Logger

/**
 * 用户修改密码
 */
func UpdatePassword(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.UpdatePassword
	if err := ctx.ValidateJSON(&p); err != nil {
		return
	}

	err := service.UpdatePassword(p.Phone, p.OldPassword, p.NewPassword)
	if err != nil {
		logger.Error("controller UpdatePassword error:", err)
		ctx.Response(common.ERROR, err.Error(), nil)
		return
	}
	// 清除session，重新登录
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	data := map[string]string{
		"data": "success",
	}
	ctx.Response(common.SUCCESS, nil, data)
}
