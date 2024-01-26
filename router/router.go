package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// IfnRegRouter is the interface for registering routers
type IfnRegRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IfnRegRouter
)

func RegRouter(fn IfnRegRouter) {
	gfnRouters = append(gfnRouters, fn)
}

// RegisterRouter is the function for registering routers
func InitRouter() {
	//优雅退出
	//ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer cancelCtx()

	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	for _, fn := range gfnRouters {
		fn(rgPublic, rgAuth)
	}

	stPort := viper.GetString("server.port")
	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("run server err: %s", err.Error()))
	}

	//<-ctx.Done()
}

// 基础平台下的路由
func InitBasePlatformRoutes() {
	InitFileRoutes()
}
