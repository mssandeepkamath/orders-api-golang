package service

import (
	"github.com/mssandeepkamaath/mediumAPI/helper"
	"github.com/mssandeepkamaath/mediumAPI/model"
	"github.com/mssandeepkamaath/mediumAPI/storage"
	"time"
)

func GetAllOrders() []model.Order {
	return storage.Orders
}

func GetOrderByID(orderId string) model.Order {
	for _, order := range storage.Orders {
		if order.OrderID == orderId {
			return order
		}
	}
	return model.Order{}
}

func AddOrder(order *model.Order) model.Order {
	order.OrderID = helper.GenerateID()
	storage.Orders = append(storage.Orders, *order)
	return *order
}

func UpdateOrder(order model.Order) bool {
	for index, o := range storage.Orders {
		if o.OrderID == order.OrderID {
			storage.Orders[index] = order
			return true
		}
	}
	return false
}

func DeleteOrder(orderId string) bool {
	for index, order := range storage.Orders {
		if order.OrderID == orderId {
			storage.Orders = append(storage.Orders[:index], storage.Orders[index+1:]...)
			return true
		}
	}
	return false
}

func DummyOrder() model.Order {
	product := model.Product{Id: "4", Name: "Samsung S10 Ultra", Price: 110000.00, Company: "Samsung"}
	address := model.Address{Flat: "#1 XYZ Apartment", Street: "2th Main Road ABC", City: "Bengaluru", State: "Karnataka", Pincode: "5600XY"}
	date := model.Date{PlacedDate: time.Now().Format("02-01-2006 15:04:05 Monday"), ExpectedShipmentDate: time.Now().Add(24 * time.Hour).Format("02-01-2006 15:04:05 Monday")}
	order := model.Order{OrderID: helper.GenerateID(), Product: product, Address: address, Date: date}
	return order
}
