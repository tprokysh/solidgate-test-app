package server

import (
	"net/http"

	"solidgate-test-app/src/config"
	"solidgate-test-app/src/db"
	"solidgate-test-app/src/server/router"

	solidgate "bitbucket.org/solidgate/go-sdk"

	customerH "solidgate-test-app/src/handlers/customer"
	customerR "solidgate-test-app/src/repositories/customer"
	customerS "solidgate-test-app/src/services/customer"

	orderH "solidgate-test-app/src/handlers/order"
	orderR "solidgate-test-app/src/repositories/order"
	orderS "solidgate-test-app/src/services/order"

	operationsH "solidgate-test-app/src/handlers/customer/operations"
	operationS "solidgate-test-app/src/services/customer/operations"

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
	dbConnection := db.GetConnection()
	apiCfg := config.GetApiConfig()
	solidgateApi := solidgate.NewSolidGateApi(apiCfg.CustomerId, apiCfg.PrivateKey, nil)

	//customer
	customerRepository := customerR.NewCustomerRepository(dbConnection)
	customerService := customerS.NewCustomerCreateService(customerRepository)
	customerHandler := customerH.NewCustomerHandler(customerService)
	//order
	orderRepository := orderR.NewOrderRepository(dbConnection)
	orderService := orderS.NewOrderCreateService(orderRepository)
	orderHandler := orderH.NewOrderHandler(orderService, solidgateApi)
	//operation
	chargeOperationService := operationS.NewChargeOperationService(orderRepository, solidgateApi)
	refundOperationService := operationS.NewRefundOperationService(orderRepository, solidgateApi)
	recurringOperationService := operationS.NewRecurringOperationService(orderRepository, solidgateApi)
	callbackOperationService := operationS.NewCallbackOperationService(orderRepository, customerRepository, solidgateApi)
	operationsHandler := operationsH.NewOperationHandler(chargeOperationService, refundOperationService, recurringOperationService, callbackOperationService)

	s.router = router.NewRouter()
	s.router.InitRoutes(orderHandler, customerHandler, operationsHandler)

	server := &http.Server{Addr: "localhost:8080", Handler: s.router.Mux}
	server.ListenAndServe()
}
