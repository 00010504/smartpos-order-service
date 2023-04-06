package repo

import "genproto/common"

type ShopI interface {
	Upsert(*common.ShopCreatedModel) error
	GetById(req *common.RequestID) (*common.ShortShop, error)
	DeleteById(*common.RequestID) error
}
