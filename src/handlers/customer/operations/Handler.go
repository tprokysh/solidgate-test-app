package operations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../../services/customer/operations"
)

type Operation struct {
	chargeService operations.Charge
}

func NewOperationHandler(service operations.Charge) Operation {
	return Operation{chargeService: service}
}

func (handler Operation) Charge(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := handler.chargeService.Charge(body)

	fmt.Println(json.Marshal(res))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
