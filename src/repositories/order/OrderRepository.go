package order

import (
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"solidgate-test-app/src/models"
)

type Order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) Order {
	return Order{db}
}

func (m *Order) Create(data []byte) ([]byte, error) {
	newOrder := &models.Order{
		Status: "pending",
	}
	json.Unmarshal(data, newOrder)

	result, _ := json.Marshal(m.db.Create(newOrder))

	return result, nil
}

func (m *Order) GetOrder(orderId string) (*models.Order, error) {
	order := &models.Order{}
	searchId, _ := strconv.Atoi(orderId)
	err := m.db.Where("id = ?", searchId).First(&order).Error

	return order, err
}

func (m *Order) UpdateOrderStatus(order *models.Order, status string) error {
	return m.db.Model(order).Update("status", status).Error
}
