package order

import (
	orderRepository "../../repositories/order"
)

type Create struct {
	repository orderRepository.Order
}

func NewOrderCreateService(repository orderRepository.Order) Create {
	return Create{repository}
}

func (service *Create) Create(data []byte) ([]byte, error) {
	newOrder, err := service.repository.Create(data)

	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}
