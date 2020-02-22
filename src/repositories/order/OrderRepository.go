package order

import (
	"strconv"

	models "../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) Order {
	return Order{db}
}

func (m *Order) Create(order *models.Order) error {
	return m.db.Create(order).Error
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
