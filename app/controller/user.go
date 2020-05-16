package controller

import (
	"encoding/json"
	"gin-user-center/app/common"
	"gin-user-center/app/config"
	"gin-user-center/app/middleware"
	"gin-user-center/app/schema"
	"gin-user-center/app/service"
	"gin-user-center/app/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var logger = common.Logger

type userSession struct {
	UserId int    `json: "userId"`
	Name   string `json: "name"`
	Avatar string `json: "avatar"`
}

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
func UploadImage(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	file, err := c.FormFile("name")
	if err != nil {
		logger.Error("controller UploadImage error:", err)
		ctx.Response(common.ERROR, common.UPDATE_AVATAR_FAIL, nil)
		return
	}

	sessionData, _ := c.Get(common.SESSION_KEY)
	var user userSession
	if err := json.Unmarshal([]byte(sessionData.(string)), &user); err != nil {
		logger.Error("controller UploadImage Unmarshal error:", err)
		ctx.Response(common.ERROR, common.SERVER_ERROR, nil)
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
	err = service.UploadImage(user.UserId, fileName)
	if err != nil {
		ctx.Response(common.ERROR, common.UPDATE_AVATAR_FAIL, nil)
		return
	}

	data := map[string]string{
		"url": config.Conf.File.UrlPrefix + fileName,
	}
	ctx.Response(common.SUCCESS, nil, data)
}

/**
 * 获取头像
 */
func GetImage(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.GetImage
	if err := ctx.ValidateRouter(&p); err != nil {
		return
	}
	fileName := config.Conf.File.DirName + p.ImageName
	c.File(fileName)
}
