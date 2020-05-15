package controller

import (
	"gin-user-center/app/common"
	"gin-user-center/app/middleware"
	"gin-user-center/app/schema"
	"gin-user-center/app/service"
	"gin-user-center/app/util"

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

/**
 * 更改头像
 */
func UploadAvatar(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	file, err := c.FormFile("name")
	if err != nil {
		logger.Error("controller UploadAvatar error:", err)
		ctx.Response(common.ERROR, common.UPDATE_AVATAR_FAIL, nil)
		return
	}
	// 暂时保存到文件，TODO:上传到oss、七牛云
	fileName, err := util.SaveToFile(file)
	if err != nil {
		logger.Error("controller SaveToFile error:", err)
		ctx.Response(common.ERROR, common.UPDATE_AVATAR_FAIL, nil)
		return
	}

	// 保存图片
	// err = service.UploadAvatar(fileName)
	if err != nil {
		ctx.Response(common.ERROR, common.UPDATE_AVATAR_FAIL, nil)
		return
	}

	data := map[string]string{
		"data": fileName,
	}
	ctx.Response(common.SUCCESS, nil, data)
}
