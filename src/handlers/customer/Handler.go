package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../models"
	"../../services/customer"
)

type Customer struct {
	service customer.Create
}

func NewCustomerHandler(service customer.Create) Customer {
	return Customer{service}
}

func (handler Customer) Create(w http.ResponseWriter, r *http.Request) {
	newCustomer := models.Customer{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.Unmarshal(body, &newCustomer)
	user, err := handler.service.Create(newCustomer)
	fmt.Println(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
