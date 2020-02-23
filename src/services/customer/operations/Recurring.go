package operations

import (
	"encoding/json"

	"../../../errors"
	orderRepository "../../../repositories/order"
	"../../../solidgate"
)

type Recurring struct {
	orderRepository orderRepository.Order
	solidgateApi    *solidgate.Api
}

func NewRecurringOperationService(orderRepository orderRepository.Order, solidgateApi *solidgate.Api) Recurring {
	return Recurring{orderRepository, solidgateApi}
}

func (service *Recurring) Recurring(data []byte) ([]byte, error) {
	reqOrder := reqOrder{}
	json.Unmarshal(data, &reqOrder)

	order, err := service.orderRepository.GetOrder(reqOrder.OrderId)
	if err != nil {
		orderError := errors.OrderError{
			Status:  "400",
			Message: "Order not found",
		}
		return json.Marshal(orderError)
	}

	res, err := service.solidgateApi.Recurring(data)
	if err != nil {
		return res, err
	}

	result := resOrder{}
	json.Unmarshal(res, &result)

	err = service.orderRepository.UpdateOrderStatus(order, result.Order.Status)
	if err != nil {
		orderError := errors.OrderError{
			Status:  "400",
			Message: "Can't update order status",
		}
		return json.Marshal(orderError)
	}

	return res, nil
}
