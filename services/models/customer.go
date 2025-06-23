package customermodel

import (
	"context"
	"database/sql"
	"fmt"

	//mariadbstore "yats/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Customer struct {
	bun.BaseModel `bun:"table:Customer,alias:customer"`
	CustomerId    int64  `bun:"customerid,pk,autoincrement,type:integer"`
	Name          string `bun:"name,notnull"`
	Email         string `bun:"email"`
	Phone         int64  `bun:"phone,type:integer"`
}

var db *bun.DB = &bun.DB{}

// need to create a a file named mariadbstore.go with the following properties for this to work.
func init() {
	sqlconstr := fmt.Sprintf("%s:%s@/test", Mariadbusername, Mariadbpassword)
	sqldb, err := sql.Open("mysql", sqlconstr)
	if err != nil {
		panic(err)
	}
	db = bun.NewDB(sqldb, mysqldialect.New())
}

// not sure if we want to return pointer to customer or dereference it.
func GetCustomerById(customerId int) *Customer {
	customer := new(Customer)
	err := db.NewSelect().Model(customer).Where("customerid = ?", customerId).Scan(context.Background())
	if err != nil {
		// return some error message. idc.
	}

	return customer
}
