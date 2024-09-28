package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetSingleProductQuantitiyByCategoryIDAndWorkingUnitID(
	dao *daos.Dao,
	categoryId string,
	workingUnitId string,
) (
	*custom_models.ProductQuantity,
	error,
) {
	var productQuantity *custom_models.ProductQuantity
	err := custom_models.ProductQuantityQuery(dao).
		Where(dbx.HashExp{"categoryId": categoryId, "workingUnitId": workingUnitId}).
		One(&productQuantity)
	if err != nil {
		return nil, err
	}
	return productQuantity, nil
}
