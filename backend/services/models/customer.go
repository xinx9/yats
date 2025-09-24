package models

import (
	"context"
	"database/sql"
	"fmt"
	mariadbstore "yats/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Customer struct {
	CustomerId int64
	Name       string
	Email      string
	Phone      int64
}
