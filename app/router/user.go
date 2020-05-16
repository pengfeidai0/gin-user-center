package router

import (
	"gin-user-center/app/controller"
	"gin-user-center/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	router := group.Group("").Use(middleware.SessionAuth())
	{
		router.POST("/change_pwd", controller.UpdatePassword)
		router.POST("/upload_avatar", controller.UploadImage)
		router.GET("/file/:imageName", controller.GetImage)
	}
}
