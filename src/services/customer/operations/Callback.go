package operations

import (
	"encoding/json"
	customerR "solidgate-test-app/src/repositories/customer"
	orderR "solidgate-test-app/src/repositories/order"

	solidgate "bitbucket.org/solidgate/go-sdk"
)

type Callback struct {
	orderRepository    orderR.Order
	customerRepository customerR.Customer
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

func NewCallbackOperationService(orderRepository orderR.Order, customerRepository customerR.Customer, solidgateApi *solidgate.Api) Callback {
	return Callback{orderRepository, customerRepository, solidgateApi}
}

func (service *Callback) Callback(data []byte) error {
	responseCard := resCard{}
	json.Unmarshal(data, &responseCard)

	if responseCard.Order.OrderId == "" || responseCard.Order.Status == "" {
		return nil
	}

	order, err := service.orderRepository.GetOrder(responseCard.Order.OrderId)
	if err != nil {
		return err
	}

	err = service.orderRepository.UpdateOrderStatus(order, responseCard.Order.Status)
	if err != nil {
		return err
	}

	customer, err := service.customerRepository.GetCustomerById(order.CustomerId)
	if err != nil {
		return err
	}

	err = service.customerRepository.UpdateCustomerToken(customer, responseCard.Transaction.Card.CardToken.Token)

	return nil
}
