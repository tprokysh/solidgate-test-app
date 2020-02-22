package order

import (
	models "../../models"
	orderRepository "../../repositories/order"
)

type Create struct {
	repository orderRepository.Order
}

func NewOrderCreateService(repository orderRepository.Order) Create {
	return Create{repository}
}

func (service *Create) Create(data models.Order) (*models.Order, error) {
	order := &models.Order{
		CustomerId: data.CustomerId,
		Status:     "pending",
	}

	err := service.repository.Create(order)

	if err != nil {
		return order, err
	}

	return order, nil
}
