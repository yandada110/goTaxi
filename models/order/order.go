package order

import (
	"goTaxi/models"
	"time"
)

const (
	// 创建订单 -
	StatusCreate = "create"
	// 司机接单 -
	StatusReceive = "receive"
	// 订单已开始 -
	StatusStart = "start"
	// 订单取消 -
	StatusCancel = "cancel"
	// 订单完成 -
	StatusComplete = "complete"
	//订单缓存地址
	PushOrderListKey = "push:order:list:key"
	//订单缓存id集合
	PushOrderListKey1 = "push:order:list1:key"
)

//var (
//	Status = []string{StatusCreate, StatusStart,StatusReceive,StatusCancel,StatusComplete}
//)
type Order struct {
	models.BaseModelId
	OrderOn         string    `gorm:"column:order_on;index:index_order_on;type:varchar(64);not null" json:"order_on"`
	PassengerUserId uint64    `gorm:"column:passenger_user_id;index:index_order_user;not null" json:"passenger_user_id"`
	DriverUserId    uint64    `gorm:"column:driver_user_id;index:index_order_user" json:"driver_user_id"`
	OrderStart      string    `gorm:"column:order_start;type:varchar(50);not null" json:"order_start"`
	OrderEnd        string    `gorm:"column:order_end;type:varchar(50);not null" json:"order_end"`
	OrderStatus     string    `gorm:"column:order_status;type:varchar(50);not null" json:"order_status"`
	OrderPrice      float64    `gorm:"column:order_price;type:decimal(10,2);default:0" json:"order_price"`
	CreatedStart    time.Time `gorm:"column:created_start;index" json:"created_start"`
	UpdateEnt       time.Time `gorm:"column:update_ent;index" json:"update_ent"`
	models.BaseModelTime
}
