package bootstrap

import (
	"goTaxi/models/order"
	"goTaxi/models/user"
	"goTaxi/pkg/database"
)

var MigrateStruct map[string]interface{}

// 初始化表结构体
func init() {
	MigrateStruct = make(map[string]interface{})
	MigrateStruct["User"] = user.User{}
	MigrateStruct["passenger"] = order.Order{}
}

func autoMigrate() {
	database.SetMysqlDB()
	for _, v := range MigrateStruct {
		_ = database.DB.AutoMigrate(v)
	}
}
