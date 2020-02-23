package server

import (
	"net/http"

	config "../config"
	database "../database"
	customerHandler "../handlers/customer"
	customerRepository "../repositories/customer"
	customerService "../services/customer"
	"../solidgate"

	orderHandler "../handlers/order"
	orderRepository "../repositories/order"
	orderService "../services/order"

	operationsHandler "../handlers/customer/operations"
	operationService "../services/customer/operations"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Server struct{}

func NewServer() (server *Server) {
	return &Server{}
}

func (s *Server) Run() {
	dbConnection := database.GetConnection()
	apiCfg := config.GetApiConfig()

	apiUrl := apiCfg.Api + apiCfg.ApiPrefix
	solidgateApi := solidgate.NewSolidGateApi(apiCfg.CustomerId, apiCfg.PrivateKey, &apiUrl)

	//TODO: move this routes into single file Router.go
	customerRepository := customerRepository.NewCustomerRepository(dbConnection)
	customerService := customerService.NewCustomerCreateService(customerRepository)
	customerHandler := customerHandler.NewCustomerHandler(customerService)

	orderRepository := orderRepository.NewOrderRepository(dbConnection)
	orderService := orderService.NewOrderCreateService(orderRepository)
	orderHandler := orderHandler.NewOrderHandler(orderService, solidgateApi)

	chargeOperationService := operationService.NewChargeOperationService(orderRepository, solidgateApi)
	refundOperationService := operationService.NewRefundOperationService(orderRepository, solidgateApi)
	callbackOperationService := operationService.NewCallbackOperationService(orderRepository, customerRepository, solidgateApi)
	operationsHandler := operationsHandler.NewOperationHandler(chargeOperationService, refundOperationService, callbackOperationService, initOperationService)

	route := mux.NewRouter()

	route.HandleFunc("/customer", customerHandler.Create).Methods("POST")
	route.HandleFunc("/order", orderHandler.Create).Methods("POST")
	route.HandleFunc("/customer/operation/charge", operationsHandler.Charge).Methods("POST")
	route.HandleFunc("/customer/operation/refund", operationsHandler.Refund).Methods("POST")
	route.HandleFunc("/callback", operationsHandler.Callback).Methods("POST")

	http.Handle("/", route)
	http.ListenAndServe("localhost:8080", nil)
}
