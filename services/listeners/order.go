package listeners

import (
	"context"
	"genproto/common"
	"genproto/report_service"

	"genproto/order_service"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/pkg/errors"
)

func (o *orderService) CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest) (*order_service.GetOrderByIDResponse, error) {

	_, err := o.strg.Shop().GetById(&common.RequestID{Id: req.ShopId, Request: req.Request})
	if err != nil {
		return nil, err
	}

	res, err := o.strg.Order().CreateOrder(req)
	if err != nil {
		return nil, err
	}

	order, err := o.strg.Order().GetOrderById(&common.RequestID{Id: res.Id, Request: req.Request})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderService) UpdateOrder(ctx context.Context, req *order_service.UpdateOrderRequest) (*common.ResponseID, error) {

	res, err := o.strg.Order().UpdateOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) CreateOrderItem(ctx context.Context, req *order_service.CreateOrderItemRequest) (*order_service.GetOrderByIDResponse, error) {

	_, err := o.strg.Order().CreateOrderItem(req)
	if err != nil {
		return nil, err
	}

	res, err := o.strg.Order().GetOrderById(&common.RequestID{Id: req.OrderId, Request: req.Request})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) UpdateOrderItem(ctx context.Context, req *order_service.UpdateOrderItemRequest) (*order_service.GetOrderByIDResponse, error) {

	updateOrderItemRes, err := o.strg.Order().UpdateOrderItem(req)
	if err != nil {
		return nil, err
	}

	res, err := o.strg.Order().GetOrderById(&common.RequestID{Id: updateOrderItemRes.Id, Request: req.Request})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) DeleteOrderItem(ctx context.Context, req *common.RequestID) (*order_service.GetOrderByIDResponse, error) {

	delRes, err := o.strg.Order().DeleteOrderItem(req)
	if err != nil {
		return nil, err
	}

	res, err := o.strg.Order().GetOrderById(&common.RequestID{Id: delRes.Id, Request: req.Request})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) UpsertOrderDiscount(ctx context.Context, req *order_service.UpsertOrderDiscountRequest) (*order_service.GetOrderByIDResponse, error) {

	tr, err := o.strg.WithTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "error while begin transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := tr.Order().UpsertOrderDiscount(req)
	if err != nil {
		return nil, errors.Wrap(err, "error while UpsertOrderDiscount")
	}

	itemsLen, err := tr.Order().GetOrderItemsLen(&common.RequestID{Id: res.Id, Request: req.Request})
	if err != nil {
		return nil, errors.Wrap(err, "error while GetOrderById")
	}

	err = tr.Order().UpdateOrderItemsDiscount(req, itemsLen)
	if err != nil {
		return nil, errors.Wrap(err, "error while UpdateOrderItemsDiscount()")
	}

	order, err := tr.Order().GetOrderById(&common.RequestID{Id: res.Id, Request: req.Request})
	if err != nil {
		return nil, errors.Wrap(err, "error while GetOrderById")
	}

	return order, nil
}

func (o *orderService) CreateOrderPays(ctx context.Context, req *order_service.CreateOrderPaysRequest) (*order_service.CreateOrderPaysResponse, error) {

	var (
		items = make([]*order_service.OrderItemCopyRequest, 0)
	)

	tr, err := o.strg.WithTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "error while begin transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := tr.Order().CreateOrderPays(req)
	if err != nil {
		return nil, err
	}

	_, err = tr.Order().UpdateOrder(&order_service.UpdateOrderRequest{
		Id:      req.OrderId,
		Request: req.Request,
		Status:  config.OrderStatusPayed,
	})
	if err != nil {
		return nil, err
	}

	order, err := tr.Order().GetOrderById(&common.RequestID{Id: req.OrderId, Request: req.Request})
	if err != nil {
		return nil, errors.Wrap(err, "error while get order")
	}

	for _, item := range order.Items {
		items = append(items, &order_service.OrderItemCopyRequest{
			Id:         item.Id,
			Price:      item.Price,
			TotalPrice: item.TotalPrice,
			ProductId:  item.ProductId,
			Value:      item.Value,
			CreatedBy:  item.CreatedBy.Id,
		})

	}

	// get cheque
	cheque, err := tr.Cheque().GetByOrderId(req.OrderId)
	if err != nil {
		return nil, errors.Wrap(err, "error while get cheque")
	}

	// make html
	htmlLink, err := o.pdf.MakeOrderHTML(&models.CreatedOrderPDFRequest{
		Id:                 order.Id,
		ExtarnalId:         order.ExternalId,
		QRCode:             req.QrCodeUrl,
		Status:             order.Status,
		TotalPirce:         order.TotalPirce,
		TotalDiscountPrice: order.Discount.Price,
		ProductSortCount:   order.ProductSortCount,
		Shop:               order.Shop,
		Client:             order.Client,
		Cashbox:            order.Cashbox,
		Company:            order.Company,
		CreatedBy:          order.CreatedBy,
		Items:              order.Items,
		Pays:               order.Pays,
		Cheque:             cheque,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while create html file for cheque")
	}

	// os.ReadFile(htmlLink)
	// make PDF
	pdfLink, err := o.pdf.HTMLtoPDF(htmlLink, req.OrderId)
	if err != nil {
		return nil, errors.Wrap(err, "error while convert html to pdf")
	}

	// create order copy inventory_service
	_, err = o.services.InventoryService().CreateOrder(ctx, &order_service.CreateOrderCopyRequest{
		OrderId:            req.OrderId,
		ExternalId:         order.ExternalId,
		Status:             order.Status.Id,
		TotalPrice:         order.TotalPirce,
		TotalDiscountPrice: order.Discount.Price,
		ProductSortCount:   order.ProductSortCount,
		Items:              items,
		ShopId:             order.Shop.Id,
		CashboxId:          order.Cashbox.Id,
		CompanyId:          order.Company.Id,
		CreatedBy:          order.CreatedBy.Id,
		Request:            req.Request,
		Client:             order.Client,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while create order copy. InventoryService")
	}

	return &order_service.CreateOrderPaysResponse{Id: res.Id, Pdf: pdfLink}, nil
}

func (o *orderService) GetOrderById(ctx context.Context, req *common.RequestID) (*order_service.GetOrderByIDResponse, error) {

	order, err := o.strg.Order().GetOrderById(req)
	if err != nil {
		return nil, err
	}

	cheque, err := o.strg.Cheque().GetByOrderId(order.Id)
	if err != nil {
		return nil, err
	}

	order.Cheque = cheque

	return order, nil
}

func (o *orderService) GetAllOrders(ctx context.Context, req *order_service.GetOrdersRequest) (*order_service.GetAllOrdersResponse, error) {

	res, err := o.strg.Order().GetAllOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) DeleteOrderById(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {

	res, err := o.strg.Order().DeleteOrderById(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderService) GetPaymentTypeAnalytics(ctx context.Context, req *report_service.GetDashboardAnalyticsReq) (*order_service.GetPaymentTypeAnalyticsResponse, error) {
	o.log.Info("GetPaymentTypeAnalytics", logger.Any("req", req))
	result, err := o.strg.Order().GetPaymentTypeAnalytics(req)
	if err != nil {
		o.log.Error("GetPaymentTypeAnalytics", logger.Any("error", err))
		return nil, err
	}
	return &order_service.GetPaymentTypeAnalyticsResponse{
		Res: result,
	}, nil
}

func (o *orderService) AddClientToOrder(ctx context.Context, in *order_service.AddClientToOrderRequest) (*order_service.GetOrderByIDResponse, error) {

	tr, err := o.strg.WithTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "error while begin transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := o.strg.Order().AddClientToOrder(in)
	if err != nil {
		return nil, err
	}

	return res, err
}
func (o *orderService) RemoveClientFromOrder(ctx context.Context, in *order_service.RemoveClientFromOrderRequest) (*order_service.GetOrderByIDResponse, error) {
	tr, err := o.strg.WithTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "error while begin transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := o.strg.Order().RemoveClientFromOrder(in)
	if err != nil {
		return nil, err
	}

	return res, err
}
