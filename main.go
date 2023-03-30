package main

import (
	"fmt"
	"gvd_project/core"
	"gvd_project/global"
	"gvd_project/router"
)

// 服务入口
func main() {
	//读取配置文件
	core.CoreConf()
	fmt.Println(global.Config)
	//设置日志格式
	global.Log = core.InitLogger()
	/*	global.Log.Warnln("测试警告日志")
		global.Log.Errorln("测试错误日志")
		global.Log.Debugln("测试debug日志")
		global.Log.Println("测试日志")*/
	//连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)

	//获取ip端口配置
	addr := global.Config.System.Addr()
	//获取路由配置
	routers := router.InitRouter()
	global.Log.Infof("gvd_project 运行在：%s", addr)
	//启动服务
	routers.Run(addr)
}
