package main

import (
	"fmt"
	"github.com/mssandeepkamaath/mediumAPI/helper"
	"github.com/mssandeepkamaath/mediumAPI/router"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	router := router.Router()
	err := http.ListenAndServe(":8080", router)
	helper.CheckNilErr(err)
}
