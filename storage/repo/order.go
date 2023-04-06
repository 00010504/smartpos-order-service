package repo

import (
	"genproto/common"
	"genproto/report_service"

	"genproto/order_service"
)

type OrderI interface {
	CreateOrder(entity *order_service.CreateOrderRequest) (*common.ResponseID, error)
	UpdateOrder(entity *order_service.UpdateOrderRequest) (*common.ResponseID, error)
	GetOrderById(*common.RequestID) (*order_service.GetOrderByIDResponse, error)
	GetAllOrders(*order_service.GetOrdersRequest) (*order_service.GetAllOrdersResponse, error)
	DeleteOrderById(*common.RequestID) (*common.ResponseID, error)
	GetOrderItemsLen(*common.RequestID) (int, error)

	// Order Item
	CreateOrderItem(*order_service.CreateOrderItemRequest) (*common.ResponseID, error)
	UpdateOrderItem(*order_service.UpdateOrderItemRequest) (*common.ResponseID, error)
	DeleteOrderItem(*common.RequestID) (*common.ResponseID, error)

	// order discount
	UpsertOrderDiscount(*order_service.UpsertOrderDiscountRequest) (*common.ResponseID, error)
	UpdateOrderItemsDiscount(discount *order_service.UpsertOrderDiscountRequest, itemsLen int) error
	// Order Pays
	CreateOrderPays(*order_service.CreateOrderPaysRequest) (*common.ResponseID, error)
	// Get Payment Types
	GetPaymentTypeAnalytics(req *report_service.GetDashboardAnalyticsReq) ([]*report_service.Payments, error)

	// client

	AddClientToOrder(in *order_service.AddClientToOrderRequest) (*order_service.GetOrderByIDResponse, error)
	RemoveClientFromOrder(in *order_service.RemoveClientFromOrderRequest) (*order_service.GetOrderByIDResponse, error)
}
