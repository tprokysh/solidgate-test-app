package customer

import (
	"io/ioutil"
	"net/http"

	"solidgate-test-app/src/services/customer"
)

type Customer struct {
	service customer.Create
}

func NewCustomerHandler(service customer.Create) Customer {
	return Customer{service}
}

func (handler Customer) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	customer, err := handler.service.Create(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(customer)
}
