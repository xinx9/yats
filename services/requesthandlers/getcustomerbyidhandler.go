package requesthandler

import (
	"context"
	dataaccess "yats/services/dataaccess"
	models "yats/services/models"
)

type GetCustomerByIdRequestHandler struct {
	dbconn dataaccess.DbContext
}

// need to set up dependency injection,
func NewGetCustomerByIdHandler(dbcontext dataaccess.DbContext) *GetCustomerByIdRequestHandler {
	return &GetCustomerByIdRequestHandler{
		dbconn: dbcontext,
	}
}

// not sure if we want to return pointer to customer or dereference it.
func (requestHandler GetCustomerByIdRequestHandler) GetCustomerById(customerId int) (*models.Customer, error) {
	customerModel := new(models.Customer)
	customerDb := new(dataaccess.Customer)
	//need to get access to bun.db to run select
	err := requestHandler.dbconn.NewSelect().Model(customerDb).Where("customerid = ?", customerId).Scan(context.Background())
	
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
