package customer

import (
	"encoding/json"
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

func (m *Customer) Create(data []byte) ([]byte, error) {
	newCustomer := &models.Customer{}
	json.Unmarshal(data, newCustomer)

	result, _ := json.Marshal(m.db.Create(newCustomer))

	return result, nil
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
