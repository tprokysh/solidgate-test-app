package order

import (
	"io/ioutil"
	"net/http"

	"bitbucket.org/solidgate/go-sdk"
	"solidgate-test-app/src/services/order"
)

type Order struct {
	service      order.Create
	solidgateApi *solidgate.Api
}

func NewOrderHandler(service order.Create, solidgateApi *solidgate.Api) Order {
	return Order{service, solidgateApi}
}

func (handler Order) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	order, err := handler.service.Create(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(order)
}
