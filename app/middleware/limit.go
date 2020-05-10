package middleware

import (
	"gin-user-center/app/config"

	"github.com/gin-gonic/gin"
)

func Limit() gin.HandlerFunc {
	sem := make(chan struct{}, config.Conf.LimitNum)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
