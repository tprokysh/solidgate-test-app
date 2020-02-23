package errors

type OrderError struct {
	Status  string
	Message string
}

func OrderNotFound() OrderError {
	orderError := OrderError{
		Status:  "400",
		Message: "Order not found",
	}

	return orderError
}

func OrderFailUpdateStatus() OrderError {
	orderError := OrderError{
		Status:  "400",
		Message: "Can't update order status",
	}

	return orderError
}

func OrderIdIsUndefined() OrderError {
	orderError := OrderError{
		Status:  "400",
		Message: "OrderId is undefined",
	}

	return orderError
}
