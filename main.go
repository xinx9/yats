package main

import (
	"fmt"
	router "yats/app/controllers"
	customermodel "yats/services/models"
)

func main() {
	//ctx := context.Background()
	// var customer = Customer_Controller.GetCustomerById(ctx, 1)
	// fmt.Printf("personId: %d, name: %s, phone: %d, email: %s ", customer.CustomerId, customer.Name, customer.Phone, customer.Email)

	var customer = customermodel.GetCustomerById(1)

	fmt.Printf("personId: %d, name: %s, phone: %d, email: %s ", customer.CustomerId, customer.Name, customer.Phone, customer.Email)
	router.SetupServer()
}
