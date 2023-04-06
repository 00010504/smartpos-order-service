package repo

import "genproto/common"

type PaymentTypeI interface {
	Upsert(entity *common.CommonPaymentTypes) error
	UpsertMany(entity []*common.CommonPaymentTypes) error
}
