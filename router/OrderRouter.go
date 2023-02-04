package router

import (
	"github.com/gorilla/mux"
	"github.com/mssandeepkamaath/mediumAPI/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/orders", controller.ListOrders).Methods("GET")
	router.HandleFunc("/api/add-order", controller.AddOrder).Methods("POST")
	router.HandleFunc("/api/order/{id}", controller.OrderById).Methods("GET")
	router.HandleFunc("/api/update-order", controller.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/delete-order/{id}", controller.DeleteOrder).Methods("DELETE")
	return router
}
