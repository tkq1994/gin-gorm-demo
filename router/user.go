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
			//rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			//rgAuthUser.POST("/userList", userApi.GetUserList)
			//rgAuthUser.PUT("/:id", userApi.UpdateUser)
			//rgAuthUser.DELETE("/:id", userApi.DeleteUser)
		}
	})
}
