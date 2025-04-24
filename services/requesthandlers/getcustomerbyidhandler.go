package requesthandler

import (
	"context"
	dataaccess "yats/services/dataaccess"
	models "yats/services/models"
	requests "yats/services/requests"

	"github.com/uptrace/bun"
)

func GetCustomerById(dbContext *bun.DB, request requests.GetCustomerByIdRequest) (*models.Customer, error) {
	customerModel := new(models.Customer)
	customerDb := new(dataaccess.Customer)
	//need to get access to bun.db to run select
	err := dbContext.NewSelect().Model(customerDb).Where("customerid = ?", request.CustomerId).Scan(context.Background())

	if err != nil {
		return nil, err
	}

	// convert db to model
	customerModel.CustomerId = customerDb.CustomerId
	customerModel.Email = customerDb.Email
	customerModel.Name = customerDb.Name
	customerModel.Phone = customerDb.Phone

	return customerModel, nil
}
