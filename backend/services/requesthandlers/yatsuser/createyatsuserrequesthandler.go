package requesthandlers

import (
	"context"
	"errors"
	dataaccess "yats/services/dataaccess"
	requests "yats/services/requests/yatsuser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func CreateYatsUserRequestHandler(request requests.CreateYatsUserRequest, db *bun.DB) (int64, error) {
	yatsuser := new(dataaccess.YatsUser)
	yatsuser.FirstName = request.FirstName
	yatsuser.LastName = request.LastName
	yatsuser.UserName = request.UserName
	res, err := db.NewInsert().Model(yatsuser).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected < 1 {
		return 0, errors.New("could not insert new YatsUser")
	}

	newRowId, _ := res.LastInsertId()

	return newRowId, nil
}
