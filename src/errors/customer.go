package errors

type CustomerError struct {
	Status  string
	Message string
}

func CustomerNotFound() CustomerError {
	customerError := CustomerError{
		Status:  "400",
		Message: "Customer not found",
	}
	return customerError
}
