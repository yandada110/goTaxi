package bootstrap

import (
	"goTaxi/config"
	"goTaxi/pkg/redis"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func Start() {
	//redis初始化
	redis.InitRedis()
	// 数据库初始化
	autoMigrate()

}