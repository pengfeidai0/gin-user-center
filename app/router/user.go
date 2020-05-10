package router

import (
	"gin-user-center/app/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	router := group.Group("")
	{
		router.POST("/user", controller.AddUser)
	}
}
