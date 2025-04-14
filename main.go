package main

import (
	"fmt"
	customermodel "yats/services/models"
)

func main() {
	var customer = customermodel.GetCustomerById(1)

	fmt.Printf("personId: %d, name: %s, phone: %d, email: %s ", customer.CustomerId, customer.Name, customer.Phone, customer.Email)
}
