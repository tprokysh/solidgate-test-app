package customer

import (
	models "../../models"
	customerRepository "../../repositories/customer"
)

type Create struct {
	repository customerRepository.Customer
}

func NewCustomerCreateService(repository customerRepository.Customer) Create {
	return Create{repository}
}

func (service *Create) Create(data models.Customer) (*models.Customer, error) {
	customer := &models.Customer{
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Email:          data.Email,
		RecurringToken: "",
	}

	err := service.repository.Create(customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}
