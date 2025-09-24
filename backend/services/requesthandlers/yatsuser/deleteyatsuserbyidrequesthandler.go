package requesthandlers

import (
	"context"
	"errors"
	dataaccess "yats/services/dataaccess"
	requests "yats/services/requests/yatsuser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func DeleteYatsUserByIdRequestHandler(request requests.DeleteYatsUserByIdRequest, db *bun.DB) (bool, error) {
	res, err := db.NewDelete().Model((*dataaccess.YatsUser)(nil)).Where("id = ?", request.Id).Exec(context.Background())
	if err != nil {
		return false, err
	}
	numDeleted, _ := res.RowsAffected()
	if numDeleted < 1 {
		return false, errors.New("User with id does not exist")
	}
	return true, nil
}
