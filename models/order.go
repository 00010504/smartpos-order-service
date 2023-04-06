package models

import (
	"database/sql"
	"genproto/common"
	"genproto/corporate_service"

	"genproto/order_service"
)

type NullOrderStatus struct {
	Id          sql.NullString
	Name        sql.NullString
	Translation map[string]string
}

type CreatedOrderPDFRequest struct {
	Id                 string
	ExtarnalId         string
	QRCode             string
	Status             *order_service.OrderStatus
	TotalPirce         float32
	TotalDiscountPrice float32
	ProductSortCount   int32
	Shop               *order_service.OrderShop
	Client             *order_service.OrderClient
	Cashbox            *order_service.OrderCashbox
	Company            *common.ShortCompany
	CreatedBy          *common.ShortUser
	Items              []*order_service.GetOrderItemResponse
	Pays               []*order_service.GetOrderPayResponse
	// Cheque             *GetChequeResponse
	Cheque             *corporate_service.Cheque
}

type OrderStatus struct {
	Id         string
	Name       string
	Tranlation map[string]string
}
