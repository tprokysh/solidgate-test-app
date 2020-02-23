package customer

import (
	customerRepository "../../repositories/customer"
)

type Create struct {
	repository customerRepository.Customer
}

func NewCustomerCreateService(repository customerRepository.Customer) Create {
	return Create{repository}
}

func (service *Create) Create(data []byte) ([]byte, error) {
	newCustomer, err := service.repository.Create(data)

	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil
}
