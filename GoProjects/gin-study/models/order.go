package models

import (
	"gorm.io/gorm"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

func GetOrderById(db *gorm.DB, id int) (*Order, error) {
	var order Order
	result := db.Where("id = ? ", id, id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	//gorm.ErrRecordNotFound

	return &order, nil
}

func UpdateOrder(db *gorm.DB, id int, count int) (order *Order, err error) {
	//db.Model(&Order{}).Where("id = ? ", id).Update("status", count)
	//db.Model(&Order{}).Updates(&Order{})
	db.Find(&order)

	//db.Clauses(dbresolver.Write)

	return order, nil
}
