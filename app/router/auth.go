package router

import (
	"gin-user-center/app/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(group *gin.RouterGroup, store sessions.Store) {
	router := group.Group("").Use(sessions.Sessions("mysession", store))
	{
		router.POST("/login", controller.Login)
	}
}
