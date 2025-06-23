package requesthandlers

import (
	"context"
	dataaccess "yats/services/dataaccess"
	models "yats/services/models"
	requests "yats/services/requests/yatsuser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func GetYatsUserById(request requests.GetYatsUserByIdRequest, db *bun.DB) (*models.YatsUser, error) {
	yatsuser := new(dataaccess.YatsUser)
	err := db.NewSelect().Model(yatsuser).Where("id = ?", request.Id).Scan(context.Background())
	if err != nil {
		return nil, err
	}

	yatsUserModel := new(models.YatsUser)
	yatsUserModel.ID = yatsuser.ID
	yatsUserModel.FirstName = yatsuser.FirstName
	yatsUserModel.LastName = yatsuser.LastName
	yatsUserModel.UserName = yatsuser.UserName

	return yatsUserModel, nil
}
