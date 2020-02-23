package operations

import (
	// "bytes"
	"encoding/json"
	"fmt"

	"../../../errors"
	customerRepository "../../../repositories/customer"
	orderRepository "../../../repositories/order"
	"../../../solidgate"
)

type Callback struct {
	orderRepository    orderRepository.Order
	customerRepository customerRepository.Customer
	solidgateApi       *solidgate.Api
}

type resCard struct {
	Transaction struct {
		Card struct {
			CardToken struct {
				Token string `json:"token"`
			} `json:"card_token"`
		} `json:"card"`
	} `json:"transaction"`
	Order struct {
		OrderId string `json:"order_id"`
		Status  string `json:"status"`
	}
}

func NewCallbackOperationService(orderRepository orderRepository.Order, customerRepository customerRepository.Customer, solidgateApi *solidgate.Api) Callback {
	return Callback{orderRepository, customerRepository, solidgateApi}
}

func (service *Callback) Callback(data []byte) ([]byte, error) {
	responseCard := resCard{}
	json.Unmarshal(data, &responseCard)

	fmt.Println("token", responseCard.Transaction.Card.CardToken.Token)
	fmt.Println("status", responseCard.Order.Status)
	fmt.Println("order_id", responseCard.Order.OrderId)

	order, err := service.orderRepository.GetOrder(responseCard.Order.OrderId)
	if err != nil {
		fmt.Println("FIRST ERROR")
		orderError := errors.OrderNotFound()
		return json.Marshal(orderError)
	}

	err = service.orderRepository.UpdateOrderStatus(order, responseCard.Order.Status)
	if err != nil {
		fmt.Println("SECOND ERROR")
		orderError := errors.OrderError{
			Status:  "400",
			Message: "Can't update order status",
		}
		return json.Marshal(orderError)
	}

	customer, err := service.customerRepository.GetCustomerById(order.CustomerId)
	if err != nil {
		fmt.Println("THIRD ERROR")
		customerError := errors.CustomerNotFound()
		return json.Marshal(customerError)
	}

	err = service.customerRepository.UpdateCustomerToken(customer, responseCard.Transaction.Card.CardToken.Token)
	res, _ := json.Marshal(responseCard)

	return res, nil
}
