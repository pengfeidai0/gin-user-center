package controller

import (
	"encoding/json"
	"gin-user-center/app/common"
	"gin-user-center/app/config"
	"gin-user-center/app/middleware"
	"gin-user-center/app/service"

	"gin-user-center/app/schema"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
	session.Set(config.Conf.Session.Key, string(value))
	session.Save()

	ctx.Response(common.SUCCESS, nil, user)
}
