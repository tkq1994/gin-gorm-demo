package router

import (
	"github.com/gin-gonic/gin"
	"iTypeService/api"
)

func InitFileRoutes() {
	RegRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.GET("/file", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "file test",
			})
		})
		rgPublic.POST("/upload", api.UploadFile)
		rgPublic.GET("/download/:id", api.DownloadFile)
	})
}
