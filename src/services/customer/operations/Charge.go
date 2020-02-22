package operations

import (
	"encoding/json"

	"../../../solidgate"
)

type Charge struct {
	solidgateApi *solidgate.Api
}

type ChargeParams struct {
	OrderId          string
	Amount           int
	Currency         string
	CardNumber       string
	CardHolder       string
	CardExpMonth     string
	CardExpYear      string
	CardCVV          string `json:"card_cvv"`
	CustomerEmail    string
	OrderDescription string
	IpAddress        string
	Platform         string
	Geo_country      string
}

func NewChargeOperationService(solidgateApi *solidgate.Api) Charge {
	return Charge{solidgateApi}
}

func (service *Charge) Charge(data []byte) ([]byte, error) {
	charge := ChargeParams{}
	json.Unmarshal(data, &charge)

	res, err := service.solidgateApi.Charge(data)

	if err != nil {
		return res, err
	}

	return res, nil
}
