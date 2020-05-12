package middleware

import (
	"gin-user-center/app/common"
	"gin-user-center/app/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context{Ctx: c}

		session := sessions.Default(c)
		if session.Get(config.Conf.Session.Key) == nil {
			ctx.Response(401, common.NOT_LOGIN, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
