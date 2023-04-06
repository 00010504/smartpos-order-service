package repo

import (
	"genproto/common"
)

type CashboxI interface {
	Upsert(entity *common.CashboxCreatedModel) error
	Delete(req *common.RequestID) (*common.ResponseID, error)
}
