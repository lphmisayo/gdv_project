package main

import (
	"fmt"
	"gvd_project/core"
	"gvd_project/global"
)

func main() {
	core.CoreConf() //读取配置文件
	fmt.Println(global.Config)
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
