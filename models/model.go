package models

import (
	"goTaxi/pkg/types"
	"time"
)

type BaseModelTime struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at;index" json:"update_at"`
}
type BaseModelId struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModelId) GetStringID() string {
	return types.Uint64ToString(a.Id)
}