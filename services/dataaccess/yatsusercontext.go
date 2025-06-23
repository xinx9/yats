package dataaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

type YatsUser struct {
	bun.BaseModel `bun:"table:yatsuser,alias:yatsuser"`
	ID            int64  `bun:"id,pk,autoincrement,notnull,type:integer"`
	FirstName     string `bun:"firstname"`
	LastName      string `bun:"lastname"`
	UserName      string `bun:"username,notnull"`
}
