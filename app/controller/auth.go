package controller

import (
	"encoding/json"
	"gin-user-center/app/common"
	"gin-user-center/app/middleware"
	"gin-user-center/app/model"
	"gin-user-center/app/service"

	"gin-user-center/app/schema"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
 * 用户注册
 */
func Register(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.Register
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

/**
 * 用户登录
 */
func Login(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	var p schema.Login
	if err := ctx.ValidateJSON(&p); err != nil {
		return
	}

	u, err := service.Login(p.Phone, p.Password)
	if err != nil {
		logger.Error("controller login error:", err)
		ctx.Response(common.ERROR, err.Error(), nil)
		return
	}

	user := map[string]interface{}{
		"userId": u.UserId,
		"name":   u.Name,
		"avatar": u.Avatar,
	}

	value, _ := json.Marshal(user)
	session := sessions.Default(c)
	session.Set(common.SESSION_KEY, string(value))
	session.Save()
	ctx.Response(common.SUCCESS, nil, user)
}

/**
 * 退出登录
 */
func Logout(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	session := sessions.Default(c)
	session.Clear()
	session.Save()
	type user struct{}

	ctx.Response(common.SUCCESS, nil, user{})
}
