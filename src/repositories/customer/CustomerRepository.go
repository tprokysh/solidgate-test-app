package customer

import (
	models "../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) Customer {
	return Customer{db}
}

func (m *Customer) Create(customer *models.Customer) error {
	return m.db.Create(customer).Error
}
