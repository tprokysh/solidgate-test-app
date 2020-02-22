package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	CustomerId string `json:"customer_id"`
	Status     string `json:"order_status"`
}
