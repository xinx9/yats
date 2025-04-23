package dataaccess

import (
	// "context"
	"database/sql"
	"fmt"
	mariadbstore "yats/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type DbContext interface {
	NewDbContext()
}

// need to create a a file named mariadbstore.go with the following properties for this to work.
func NewDbContext() *bun.DB {
	sqlconstr := fmt.Sprintf("%s:%s@/test", mariadbstore.Mariadbusername, mariadbstore.Mariadbpassword)
	sqldb, err := sql.Open("mysql", sqlconstr)
	if err != nil {
		panic(err)
	}
	return bun.NewDB(sqldb, mysqldialect.New())
}
