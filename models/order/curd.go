package order

import "goTaxi/pkg/database"

func (o *Order) Create() error {
	return database.DB.Create(&o).Error
}

