package customer

import (
	"strconv"

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

func (m *Customer) GetCustomerById(customerId string) (*models.Customer, error) {
	customer := &models.Customer{}
	searchId, _ := strconv.Atoi(customerId)
	err := m.db.Where("id = ?", searchId).First(&customer).Error

	return customer, err
}

func (m *Customer) UpdateCustomerToken(customer *models.Customer, token string) error {
	if customer.RecurringToken != "" {
		return nil
	}
	return m.db.Model(customer).Update("recurring_token", token).Error
}
