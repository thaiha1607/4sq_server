package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

// Return a single product quantity summary by category ID
func GetSingleProductQuantityByCategoryIDAndWorkingUnitID(
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

// Return a list of product quantities by category ID sorted by priority DESC, qty DESC
func GetProductQuantitiesByCategoryID(
	dao *daos.Dao,
	categoryId string,
) (
	[]*custom_models.ProductQuantity,
	error,
) {
	var productQuantities []*custom_models.ProductQuantity
	err := custom_models.ProductQuantityQuery(dao).
		Where(dbx.HashExp{"categoryId": categoryId}).
		OrderBy("priority DESC", "qty DESC").
		All(&productQuantities)
	if err != nil {
		return nil, err
	}
	return productQuantities, nil
}
