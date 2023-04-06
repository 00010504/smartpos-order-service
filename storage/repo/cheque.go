package repo

import (
	"genproto/common"
	"genproto/corporate_service"

	"github.com/Invan2/invan_order_service/models"
)

type ChequeI interface {
	Upsert(entity *common.ChequeCopyRequest) error
	UpsertMany(entity []*common.ChequeCopyRequest) error
	GetById(entity *common.RequestID) (*models.GetChequeResponse, error)
	GetByOrderId(orderId string) (*corporate_service.Cheque, error)
}
