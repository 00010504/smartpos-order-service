package repo

import (
	"genproto/common"
	"genproto/marketing_service"
)

type ClientI interface {
	Upsert(*marketing_service.ShortClient) error
	Delete(*common.RequestID) error
}
