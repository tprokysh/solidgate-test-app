package router

import (
	"github.com/gorilla/mux"
	"solidgate-test-app/src/handlers/customer"
	"solidgate-test-app/src/handlers/customer/operations"
	"solidgate-test-app/src/handlers/order"
)

type Router struct {
	Mux *mux.Router
}

func NewRouter() Router {
	return Router{Mux: mux.NewRouter()}
}

func (r *Router) InitRoutes(orderHandler order.Order, customerHandler customer.Customer, operationsHandler operations.Operation) {
	r.Mux.HandleFunc(
		"/customer",
		customerHandler.Create).
		Methods("POST")

	r.Mux.HandleFunc(
		"/order",
		orderHandler.Create).
		Methods("POST")

	r.Mux.HandleFunc(
		"/customer/operation/charge",
		operationsHandler.Charge).
		Methods("POST")

	r.Mux.HandleFunc(
		"/customer/operation/refund",
		operationsHandler.Refund).
		Methods("POST")

	r.Mux.HandleFunc(
		"/customer/operation/recurring",
		operationsHandler.Recurring).
		Methods("POST")

	r.Mux.HandleFunc(
		"/callback",
		operationsHandler.Callback).
		Methods("POST")
}
