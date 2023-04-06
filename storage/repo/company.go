package repo

import (
	"genproto/common"
)

type CompanyI interface {
	Upsert(entity *common.CompanyCreatedModel) error
	Delete(req *common.RequestID) (*common.ResponseID, error)
}
