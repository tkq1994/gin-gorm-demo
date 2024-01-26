package cmd

import (
	"fmt"
	"iTypeService/conf"
	"iTypeService/router"
)

// Start is the entry point of the program
func Start() {
	// 初始化配置文件
	conf.InitConfig()
	// 初始化路由
	router.InitRouter()
}

// Clean is the cleanup function
func Clean() {
	fmt.Println("==========clean=========")
}
