package main

import (
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20200223200703(txn *gorm.DB) {
	type Order struct {
		gorm.Model
		CustomerId string
		Status     string
	}
	txn.CreateTable(&Order{})
}

// Down is executed when this migration is rolled back
func Down_20200223200703(txn *gorm.DB) {
	txn.DropTable("orders")
}
