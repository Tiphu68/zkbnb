package failtx

import (
	"gorm.io/gorm"

	table "github.com/bnb-chain/zkbas/common/model/tx"
	"github.com/bnb-chain/zkbas/errorcode"
	"github.com/bnb-chain/zkbas/pkg/multcache"
)

type model struct {
	table string
	db    *gorm.DB
	cache multcache.MultCache
}

/*
	Func: CreateFailTx
	Params: failTx *FailTx
	Return: err error
	Description: create fail txVerification
*/
func (m *model) CreateFailTx(failTx *table.FailTx) error {
	dbTx := m.db.Table(m.table).Create(failTx)
	if dbTx.Error != nil {
		return errorcode.DbErrSqlOperation
	}
	if dbTx.RowsAffected == 0 {
		return errorcode.DbErrFailToCreateFailTx
	}
	return nil
}
