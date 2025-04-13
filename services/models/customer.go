package customermodel

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func initDb() {
	sqldb, err := sql.Open("mysql", "")

	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
}
