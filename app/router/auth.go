package router

import (
	"gin-user-center/app/controller"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(group *gin.RouterGroup) {
	router := group.Group("")
	// .Use(sessions.Sessions("mysession", store))
	{
		router.POST("/register", controller.Register)
		router.POST("/login", controller.Login)
		router.POST("/logout", controller.Logout)
	}
}
