package operations

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"../../../services/customer/operations"
)

type Operation struct {
	chargeService operations.Charge
}

func NewOperationHandler(chargeService operations.Charge) Operation {
	return Operation{chargeService}
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
