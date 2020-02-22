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
