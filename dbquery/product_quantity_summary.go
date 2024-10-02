package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetSingleProductQuantitySummary(
	dao *daos.Dao,
	id string,
) (
	*custom_models.ProductQuantitySummary,
	error,
) {
	var productQuantitySummary *custom_models.ProductQuantitySummary
	err := custom_models.ProductQuantitySummaryQuery(dao).
		Where(dbx.HashExp{"id": id}).
		One(&productQuantitySummary)
	if err != nil {
		return nil, err
	}
	return productQuantitySummary, nil
}

func GetBatchProductQuantitySummaries(
	dao *daos.Dao,
	ids []string,
) (
	[]*custom_models.ProductQuantitySummary,
	error,
) {
	var productQuantitySummaries []*custom_models.ProductQuantitySummary
	query := custom_models.ProductQuantitySummaryQuery(dao)
	for _, id := range ids {
		query = query.OrWhere(dbx.HashExp{"id": id})
	}
	err := query.
		All(&productQuantitySummaries)
	if err != nil {
		return nil, err
	}
	return productQuantitySummaries, nil
}

func GetAllProductQuantitySummaries(
	dao *daos.Dao,
) (
	[]*custom_models.ProductQuantitySummary,
	error,
) {
	var productQuantitySummaries []*custom_models.ProductQuantitySummary
	err := custom_models.ProductQuantitySummaryQuery(dao).
		All(&productQuantitySummaries)
	if err != nil {
		return nil, err
	}
	return productQuantitySummaries, nil
}
