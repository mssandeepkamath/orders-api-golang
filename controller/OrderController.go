package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mssandeepkamaath/mediumAPI/helper"
	"github.com/mssandeepkamaath/mediumAPI/model"
	"github.com/mssandeepkamaath/mediumAPI/service"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server is Live</h1> <p> Please refer README.md of <b>https://github.com/mssandeepkamath</b> for end-points</p><p><b>Request Body Example:</b></p>"))
	order := service.DummyOrder()
	byteDate, err := json.MarshalIndent(order, "", "\t")
	helper.CheckNilErr(err)
	w.Write(byteDate)
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := service.GetAllOrders()
	json.NewEncoder(w).Encode(orders)
}

func OrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	order := service.GetOrderByID(params["id"])
	if helper.IsOrderValid(&order) {
		json.NewEncoder(w).Encode(order)
	} else {
		json.NewEncoder(w).Encode("Order with order_id: " + params["id"] + " not found")
	}
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	helper.CheckNilErr(err)
	order = service.AddOrder(&order)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	helper.CheckNilErr(err)
	if service.UpdateOrder(order) {
		json.NewEncoder(w).Encode(order)
	} else {
		json.NewEncoder(w).Encode("Order with order_id: " + order.OrderID + " not found")
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if service.DeleteOrder(params["id"]) {
		json.NewEncoder(w).Encode("successful")
	} else {
		json.NewEncoder(w).Encode("Order with order_id: " + params["id"] + " not found")
	}
}
