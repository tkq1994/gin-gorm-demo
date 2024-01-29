package router

import (
	"github.com/gin-gonic/gin"
	"iTypeService/api"
)

func InitUserRoutes() {
	RegRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()

		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("/:id", userApi.GetUserById)
		}
	})
}
