package operations

import (
	// "fmt"
	"io/ioutil"
	"net/http"

	"../../../services/customer/operations"
)

type Operation struct {
	chargeService   operations.Charge
	refundService   operations.Refund
	callbackService operations.Callback
}

func NewOperationHandler(chargeService operations.Charge, refundService operations.Refund, callbackService operations.Callback) Operation {
	return Operation{chargeService, refundService, callbackService}
}

func (handler Operation) Charge(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := handler.chargeService.Charge(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (handler Operation) Refund(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := handler.refundService.Refund(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (handler Operation) Callback(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := handler.callbackService.Callback(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
