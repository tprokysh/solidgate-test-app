package customer

import (
	models "../../models"
	customerRepository "../../repositories/customer"
)

type Get struct {
	repository customerRepository.Customer
}

func NewCustomerGetService(repository customerRepository.Customer) Get {
	return Get{repository}
}

func (service *Get) Get(customerId string) (*models.Customer, error) {
	customer, err := service.repository.GetCustomerById(customerId)

	if err != nil {
		return customer, err
	}

	return customer, nil
}
