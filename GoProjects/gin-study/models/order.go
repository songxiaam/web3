package models

import (
	"gin-study/database"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

func GetOrderById(id int) (*Order, error) {
	var order Order
	result := database.PtDB.Where("id = ? ", id, id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
