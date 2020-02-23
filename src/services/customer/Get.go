package customer

import (
	"solidgate-test-app/src/models"
	customerR "solidgate-test-app/src/repositories/customer"
)

type Get struct {
	repository customerR.Customer
}

func NewCustomerGetService(repository customerR.Customer) Get {
	return Get{repository}
}

func (service *Get) Get(customerId string) (*models.Customer, error) {
	customer, err := service.repository.GetCustomerById(customerId)

	if err != nil {
		return customer, err
	}

	return customer, nil
}
