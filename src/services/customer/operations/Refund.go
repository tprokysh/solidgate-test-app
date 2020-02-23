package operations

import (
	"encoding/json"

	"../../../errors"
	orderRepository "../../../repositories/order"
	"../../../solidgate"
)

type Refund struct {
	orderRepository orderRepository.Order
	solidgateApi    *solidgate.Api
}

func NewRefundOperationService(orderRepository orderRepository.Order, solidgateApi *solidgate.Api) Refund {
	return Refund{orderRepository, solidgateApi}
}

func (service *Refund) Refund(data []byte) ([]byte, error) {
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

	res, err := service.solidgateApi.Refund(data)
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
