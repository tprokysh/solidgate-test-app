package main

import (
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20200223200909(txn *gorm.DB) {
	type Customer struct {
		gorm.Model
		FirstName      string
		LastName       string
		Email          string
		RecurringToken string
	}

	txn.CreateTable(&Customer{})
}

// Down is executed when this migration is rolled back
func Down_20200223200909(txn *gorm.DB) {
	txn.DropTable("customers")
}
