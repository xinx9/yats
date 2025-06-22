package controllers

import (
	"context"
	"database/sql"
	"fmt"
	mariadbstore "yats/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Customers struct {
	bun.BaseModel `bun:"table:Customers,alias:customers"`
	CustomerId    int64  `bun:"customerid,pk,autoincrement,type:integer"`
	Name          string `bun:"name,notnull"`
	Email         string `bun:"email"`
	Phone         int64  `bun:"phone,type:integer"`
}

var db *bun.DB = &bun.DB{}

// need to create a a file named mariadbstore.go with the following properties for this to work.
func init() {
	sqlconstr := fmt.Sprintf("%s:%s@/test", mariadbstore.Mariadbusername, mariadbstore.Mariadbpassword)
	sqldb, err := sql.Open("mysql", sqlconstr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("200 OK")
	}
	db = bun.NewDB(sqldb, mysqldialect.New())
}

// not sure if we want to return pointer to customer or dereference it.
func GetCustomerById(ctx context.Context, customerId int) *Customers {
	customers := new(Customers)
	err := db.NewSelect().Model(customers).Where("customerid = ?", customerId).Scan(ctx)
	if err != nil {
		panic(err)
	}

	return customers
}

func GetAllCustomers(ctx context.Context) []*Customers {
	customers := make([]*Customers, 0)
	err := db.NewSelect().Model(&customers).Scan(ctx)
	if err != nil {
		panic(err)
	}
	return customers
}
