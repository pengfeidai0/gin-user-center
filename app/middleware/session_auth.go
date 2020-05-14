package middleware

import (
	"gin-user-center/app/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context{Ctx: c}

		session := sessions.Default(c)
		if session.Get(common.SESSION_KEY) == nil {
			ctx.Response(401, common.NOT_LOGIN, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
