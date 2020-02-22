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

	customerRepository := customerRepository.NewCustomerRepository(dbConnection)
	customerService := customerService.NewCustomerCreateService(customerRepository)
	customerHandler := customerHandler.NewCustomerHandler(customerService)

	orderRepository := orderRepository.NewOrderRepository(dbConnection)
	orderService := orderService.NewOrderCreateService(orderRepository)
	orderHandler := orderHandler.NewOrderHandler(orderService, solidgateApi)

	route := mux.NewRouter()

	route.HandleFunc("/customer", customerHandler.Create).Methods("POST")
	route.HandleFunc("/order", orderHandler.Create).Methods("POST")

	http.Handle("/", route)
	http.ListenAndServe("localhost:8080", nil)
}
