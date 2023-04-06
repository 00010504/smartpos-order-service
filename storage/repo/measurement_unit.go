package repo

import (
	"genproto/common"
)

type MeasurementI interface {
	Upsert(entity *common.MeasurementUnitCopyRequest) error
	UpsertMany(entity *common.MeasurementUnitsCopyRequest) error
	Delete(req *common.RequestID) (*common.ResponseID, error)
}
