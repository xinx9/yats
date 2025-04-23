package dataaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"table:Customer,alias:customer"`
	CustomerId    int64  `bun:"customerid,pk,autoincrement,type:integer"`
	Name          string `bun:"name,notnull"`
	Email         string `bun:"email"`
	Phone         int64  `bun:"phone,type:integer"`
}
