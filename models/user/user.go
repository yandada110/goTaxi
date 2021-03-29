package user

import (
	"goTaxi/models"
)

const (
	TypePassenger = "passenger" //乘客
	TypeDriver    = "driver"    //司机
	OVERTIME         = 60                //超时时间60分钟
	UserPassengerKey = "user:passenger:" //乘客用户信息缓存
	UserDriverKey    = "user:Driver:"    //司机用户信息缓存
)

type User struct {
	models.BaseModelId
	UserName string `gorm:"column:username;type:varchar(50);not null" json:"username"`
	Phone    string `gorm:"column:phone;index:index_phone;type:varchar(11);not null" json:"phone"`
	Password string `gorm:"column:password;type:varchar(64);not null" json:"-"`
	Types    string `gorm:"column:type;type:varchar(20);not null" json:"type"`
	models.BaseModelTime
}
