package operations

import (
	"encoding/json"

	"solidgate-test-app/src/errors"
	orderR "solidgate-test-app/src/repositories/order"

	solidgate "bitbucket.org/solidgate/go-sdk"
)

type Charge struct {
	orderRepository orderR.Order
	solidgateApi    *solidgate.Api
}

type resOrder struct {
	Order struct {
		Status string
	}
}

type reqOrder struct {
	OrderId string
}

func NewChargeOperationService(orderRepository orderR.Order, solidgateApi *solidgate.Api) Charge {
	return Charge{orderRepository, solidgateApi}
}

func (service *Charge) Charge(data []byte) ([]byte, error) {
	reqOrder := reqOrder{}
	json.Unmarshal(data, &reqOrder)

	order, err := service.orderRepository.GetOrder(reqOrder.OrderId)
	if err != nil {
		orderError := errors.OrderNotFound()

		return json.Marshal(orderError)
	}

	res, err := service.solidgateApi.Charge(data)

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
