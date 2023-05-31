package main

import (
	"fmt"
	"furumvv2/controller"
	"furumvv2/dao/mysql"
	"furumvv2/logger"
	"furumvv2/pkg/snowflake"
	"furumvv2/router"
	"furumvv2/setting"
	"go.uber.org/zap"
)

func main() {
	//1.加载配置——（远程/配置文件中）
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed,err:%#v\n", err)
		return
	}

	//2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%#v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success。。。")

	//3.初始化Mysql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%#v\n", err)
		return
	}

	//初始化雪花算法ID生成器
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%#v\n", err)
		return
	}

	//初始化gin框架内置的校验器使用的翻译器——因为validator库自带的错误描述不易于读
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init transfer failed,err:%#v\n", err)
		return
	}

	//初始化引擎
	r := router.SetUpRouter(setting.Conf.Mode)

	//启动
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
