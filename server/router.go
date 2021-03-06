package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heroku/go-with-me-app/handler"
	"github.com/heroku/go-with-me-app/service"
)

func Router(services *service.Services) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/v1/booking/create", handler.CreateBookingHandler(services)).Methods("POST")
	router.HandleFunc("/v1/booking/status", handler.GetBookingStatusHandler(services)).Methods("GET")
	router.HandleFunc("/v1/booking/status/driver_not_found", handler.DriverNotFoundHandler(services)).Methods("PUT")
	router.HandleFunc("/v1/booking/fare/estimate", handler.EstimateFareHandler(services)).Methods("GET")
	router.HandleFunc("/v1/customer/create", handler.CreateCustomerHandler(services)).Methods("POST")
	router.HandleFunc("/v1/driver/create", handler.CreateDriverHandler(services)).Methods("POST")
	router.HandleFunc("/v1/driver/status", handler.DriverStatusUpdateHandler(services)).Methods("PUT")
	router.HandleFunc("/v1/driver/status", handler.GetDriverStatusHandler(services)).Methods("GET")
	router.HandleFunc("/v1/trip/start", handler.StartTripHandler(services)).Methods("PUT")
	router.HandleFunc("/v1/trip/complete", handler.CompleteTripHandler(services)).Methods("PUT")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	return router
}
