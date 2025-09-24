package requesthandlers

import (
	"context"
	"errors"
	dataaccess "yats/services/dataaccess"
	requests "yats/services/requests/yatsuser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func UpdateYatsUserRequestHandler(request requests.UpdateYatsUserRequest, db *bun.DB) (bool, error) {
	yatsuser := new(dataaccess.YatsUser)
	yatsuser.ID = request.ID
	yatsuser.FirstName = request.FirstName
	yatsuser.LastName = request.LastName
	yatsuser.UserName = request.UserName

	res, err := db.NewUpdate().Model(yatsuser).WherePK().Exec(context.Background())
	if err != nil {
		return false, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected < 1 {
		return false, errors.New("Could not update the user.")
	} else if rowsAffected > 1 {
		return false, errors.New("Multiple entries were updated when one was expected.")
	}

	return true, nil
}
