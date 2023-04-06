package listeners

import (
	"context"
	"genproto/common"
	"genproto/report_service"

	"genproto/order_service"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/events"
	"github.com/Invan2/invan_order_service/pkg/logger"
	pdfmaker "github.com/Invan2/invan_order_service/pkg/pdf"
	"github.com/Invan2/invan_order_service/services"

	"github.com/Invan2/invan_order_service/storage"
)

type orderService struct {
	log      logger.Logger
	kafka    events.PubSubServer
	strg     storage.StorageI
	pdf      pdfmaker.PdFMakekerI
	services services.GRPCServices
}

type OrderService interface {
	Ping(ctx context.Context, message *common.PingPong) (*common.PingPong, error)

	// order
	CreateOrder(context.Context, *order_service.CreateOrderRequest) (*order_service.GetOrderByIDResponse, error)
	UpdateOrder(context.Context, *order_service.UpdateOrderRequest) (*common.ResponseID, error)
	GetOrderById(context.Context, *common.RequestID) (*order_service.GetOrderByIDResponse, error)
	GetAllOrders(context.Context, *order_service.GetOrdersRequest) (*order_service.GetAllOrdersResponse, error)
	DeleteOrderById(context.Context, *common.RequestID) (*common.ResponseID, error)
	// Order items
	CreateOrderItem(context.Context, *order_service.CreateOrderItemRequest) (*order_service.GetOrderByIDResponse, error)

	UpdateOrderItem(context.Context, *order_service.UpdateOrderItemRequest) (*order_service.GetOrderByIDResponse, error)
	DeleteOrderItem(context.Context, *common.RequestID) (*order_service.GetOrderByIDResponse, error)
	// order discount
	UpsertOrderDiscount(ctx context.Context, req *order_service.UpsertOrderDiscountRequest) (*order_service.GetOrderByIDResponse, error)
	// Order pays
	CreateOrderPays(context.Context, *order_service.CreateOrderPaysRequest) (*order_service.CreateOrderPaysResponse, error)
	GetPaymentTypeAnalytics(ctx context.Context, req *report_service.GetDashboardAnalyticsReq) (*order_service.GetPaymentTypeAnalyticsResponse, error)

	// client
	AddClientToOrder(ctx context.Context, in *order_service.AddClientToOrderRequest) (*order_service.GetOrderByIDResponse, error)
	RemoveClientFromOrder(ctx context.Context, in *order_service.RemoveClientFromOrderRequest) (*order_service.GetOrderByIDResponse, error)
}

func NewOrderService(log logger.Logger, kafka events.PubSubServer, strg storage.StorageI, pdf pdfmaker.PdFMakekerI, cfg *config.Config) (OrderService, error) {

	services, err := services.NewGrpcServices(log, cfg)
	if err != nil {
		return nil, err
	}

	return &orderService{
		log:      log,
		kafka:    kafka,
		strg:     strg,
		pdf:      pdf,
		services: services,
	}, nil
}
