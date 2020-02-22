package order

import (
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
