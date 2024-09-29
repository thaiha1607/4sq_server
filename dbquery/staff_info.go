package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetStaffsByRole(dao *daos.Dao, role string) ([]*custom_models.StaffInfo, error) {
	var staffs []*custom_models.StaffInfo
	err := custom_models.StaffInfoQuery(dao).
		Where(dbx.HashExp{"role": role}).
		All(&staffs)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}

func GetStaffsByWorkingUnitId(dao *daos.Dao, workingUnitId string) ([]*custom_models.StaffInfo, error) {
	var staffs []*custom_models.StaffInfo
	err := custom_models.StaffInfoQuery(dao).
		Where(dbx.HashExp{"workingUnitId": workingUnitId}).
		All(&staffs)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}
