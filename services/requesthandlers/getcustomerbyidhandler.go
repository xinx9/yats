package requesthandler

import (
	"context"
	dataaccess "yats/services/dataaccess"
	models "yats/services/models"

	"github.com/uptrace/bun"
)

type GetCustomerByIdRequestHandler struct {
	db    *bun.DB
	query GetCustomerByIdRequest
}

func (requestHandler GetCustomerByIdRequestHandler) GetCustomerById(dbContext *bun.DB, customerId int) (*models.Customer, error) {
	customerModel := new(models.Customer)
	customerDb := new(dataaccess.Customer)
	//need to get access to bun.db to run select
	err := requestHandler.db.NewSelect().Model(customerDb).Where("customerid = ?", customerId).Scan(context.Background())

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
