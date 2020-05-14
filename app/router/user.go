package router

import (
	"gin-user-center/app/controller"
	"gin-user-center/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	router := group.Group("").Use(middleware.SessionAuth())
	{
		router.POST("/user", controller.AddUser)
	}
}
