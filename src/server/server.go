package server

import (
	"net/http"

	"./router"

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

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Server struct {
	router router.Router
}

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
	recurringOperationService := operationService.NewRecurringOperationService(orderRepository, solidgateApi)
	callbackOperationService := operationService.NewCallbackOperationService(orderRepository, customerRepository, solidgateApi)
	operationsHandler := operationsHandler.NewOperationHandler(chargeOperationService, refundOperationService, recurringOperationService, callbackOperationService)

	s.router = router.NewRouter()
	s.router.InitRoutes(orderHandler, customerHandler, operationsHandler)

	server := &http.Server{Addr: "localhost:8080", Handler: s.router.Mux}
	server.ListenAndServe()
}
