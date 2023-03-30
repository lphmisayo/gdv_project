package main

import (
	"fmt"
	"gvd_project/core"
	"gvd_project/global"
)

// 服务入口
func main() {
	core.CoreConf() //读取配置文件
	fmt.Println(global.Config)
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
