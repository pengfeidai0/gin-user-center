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
		sessionData := session.Get(common.SESSION_KEY)
		if sessionData == nil {
			ctx.Response(401, common.NOT_LOGIN, nil)
			c.Abort()
			return
		}
		c.Set(common.SESSION_KEY, sessionData)
		c.Next()
	}
}
