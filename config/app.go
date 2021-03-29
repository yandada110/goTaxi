package config

import (
	"goTaxi/pkg/config"
	"goTaxi/pkg/redis"
)

// Initialize 配置信息初始化
func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法
	redis.InitRedis()
}
func init() {
	config.Add("redis",config.StrMap{
		"addr":config.Env("RDB_ADDR","localhost:6379"),
		"password":config.Env("RDB_PASSWORD",""),
		"db":config.Env("RDB_DB",1),
	})
}

