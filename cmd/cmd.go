package cmd

import (
	"fmt"
	"iTypeService/conf"
	"iTypeService/global"
	"iTypeService/router"
	"iTypeService/utils"
)

// Start is the entry point of the program
func Start() {
	var initErr error
	// 初始化配置文件
	conf.InitConfig()
	//	初始化日志组件
	global.Logger = conf.InitLogger()
	//	初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	// 初始化路由
	router.InitRouter()
}

// Clean is the cleanup function
func Clean() {
	fmt.Println("==========clean=========")
}
