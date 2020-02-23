package operations

import (
	"encoding/json"

	"solidgate-test-app/src/errors"
	orderR "solidgate-test-app/src/repositories/order"

	solidgate "bitbucket.org/solidgate/go-sdk"
)

type Refund struct {
	orderRepository orderR.Order
	solidgateApi    *solidgate.Api
}

func NewRefundOperationService(orderRepository orderR.Order, solidgateApi *solidgate.Api) Refund {
	return Refund{orderRepository, solidgateApi}
}

func (service *Refund) Refund(data []byte) ([]byte, error) {
	reqOrder := reqOrder{}
	json.Unmarshal(data, &reqOrder)

	order, err := service.orderRepository.GetOrder(reqOrder.OrderId)
	if err != nil {
		orderError := errors.OrderNotFound()

		return json.Marshal(orderError)
	}

	res, err := service.solidgateApi.Refund(data)
	if err != nil {
		return res, err
	}

	result := resOrder{}
	json.Unmarshal(res, &result)

	if result.Order.Status == "" {
		return res, err
	}

	err = service.orderRepository.UpdateOrderStatus(order, result.Order.Status)
	if err != nil {
		orderError := errors.OrderFailUpdateStatus()

		return json.Marshal(orderError)
	}

	return res, nil
}
