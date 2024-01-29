package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"iTypeService/global"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IfnRegRouter is the interface for registering routers
type IfnRegRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IfnRegRouter
)

func RegRouter(fn IfnRegRouter) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

// RegisterRouter is the function for registering routers
func InitRouter() {
	//创建监听ctrl + c，应用退出信号的上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	//初始化gin框架，并注册相关路由
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	//初始基础平台的路由
	InitBasePlatformRoutes()

	//开始注册系统各模块对应的路由信息
	for _, fn := range gfnRouters {
		fn(rgPublic, rgAuth)
	}

	//继承swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8083"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("run server err: %s", err.Error()))
	}

	<-ctx.Done()

	ctx, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("stop server error %s", err.Error()))
	}

	global.Logger.Info(fmt.Sprintf("stop server success"))
}

// 基础平台下的路由
func InitBasePlatformRoutes() {
	InitFileRoutes()
	InitUserRoutes()
}
