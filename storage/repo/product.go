package repo

import (
	"genproto/catalog_service"
	"genproto/common"
)

type ProductI interface {
	Upsert(entity *common.CreateProductCopyRequest) error
	Delete(req *common.RequestID) (*common.ResponseID, error)
	UpsertShopMeasurmentValue(req *catalog_service.UpsertShopMeasurmentValueRequest) error
	InsertMany([]*common.CreateProductCopyRequest) error
	BulkUpdate(req *catalog_service.ProductBulkOperationRequest) (*common.ResponseID, error)
}
