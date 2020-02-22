package operations

import (
	"encoding/json"
	"fmt"

	"../../../errors"
	orderRepository "../../../repositories/order"
	"../../../solidgate"
)

type Charge struct {
	orderRepository orderRepository.Order
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

func NewChargeOperationService(orderRepository orderRepository.Order, solidgateApi *solidgate.Api) Charge {
	return Charge{orderRepository, solidgateApi}
}

func (service *Charge) Charge(data []byte) ([]byte, error) {
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

	res, err := service.solidgateApi.Charge(data)
	if err != nil {
		return res, err
	}

	result := resOrder{}
	json.Unmarshal(res, &result)

	fmt.Println("status", result.Order.Status)

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
