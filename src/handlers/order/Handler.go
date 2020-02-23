package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../models"
	"../../services/order"
	"../../solidgate"
)

type Order struct {
	service      order.Create
	solidgateApi *solidgate.Api
}

func NewOrderHandler(service order.Create, solidgateApi *solidgate.Api) Order {
	return Order{service, solidgateApi}
}

func (handler Order) Create(w http.ResponseWriter, r *http.Request) {
	newOrder := models.Order{}
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newOrder)

	order, err := handler.service.Create(newOrder)

	fmt.Println(order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
