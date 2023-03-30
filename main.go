package main

import (
	"fmt"
	"gvd_project/core"
	"gvd_project/global"
)

// 服务入口
func main() {
	//读取配置文件
	core.CoreConf()
	fmt.Println(global.Config)
	//设置日志格式
	global.Log = core.InitLogger()
	global.Log.Warnln("测试警告日志")
	global.Log.Errorln("测试错误日志")
	global.Log.Debugln("测试debug日志")
	global.Log.Println("测试日志")
	//连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
