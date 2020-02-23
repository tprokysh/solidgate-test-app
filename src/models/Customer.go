package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	RecurringToken string `json:"token"`
}
