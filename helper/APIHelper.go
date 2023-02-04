package helper

import (
	"github.com/mssandeepkamaath/mediumAPI/model"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsOrderValid(order *model.Order) bool {
	if order.OrderID != "" {
		return true
	}
	return false
}

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(math.MaxInt64))
}
