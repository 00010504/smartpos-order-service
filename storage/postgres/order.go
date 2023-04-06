package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"genproto/catalog_service"
	"genproto/common"
	"genproto/report_service"
	"strings"

	"genproto/order_service"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/helper"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type orderRepo struct {
	db  models.DB
	log logger.Logger
}

func NewOrderRepo(log logger.Logger, db models.DB) repo.OrderI {
	return &orderRepo{
		db:  db,
		log: log,
	}
}

func (o *orderRepo) CreateOrder(entity *order_service.CreateOrderRequest) (*common.ResponseID, error) {

	orderId := uuid.New().String()

	query := `
		INSERT INTO
			"order"
		(
			id,
			shop_id,
			cashbox_id,
			company_id,
			cashier_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := o.db.Exec(
		query,
		orderId,
		entity.ShopId,
		entity.CashboxId,
		entity.Request.CompanyId,
		entity.Request.UserId,
		entity.Request.UserId,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while create order")
	}

	return &common.ResponseID{Id: orderId}, nil
}

func (o *orderRepo) UpdateOrder(entity *order_service.UpdateOrderRequest) (*common.ResponseID, error) {

	query := `
		UPDATE
			"order"
		SET
			status = $2
		WHERE
			id = $1 AND deleted_at = 0
	`

	res, err := o.db.Exec(query, entity.Id, entity.Status)
	if err != nil {
		return nil, errors.Wrap(err, "error while update order")
	}

	i, err := res.RowsAffected()
	if err != nil {
		return nil, errors.New("order item not found. RowsAffected")
	}

	if i == 0 {
		return nil, errors.New("Order not found")
	}

	return &common.ResponseID{Id: entity.Id}, nil
}

func (o *orderRepo) GetOrderItemsLen(req *common.RequestID) (int, error) {

	var res int

	query := `
		SELECT count(*)
		FROM "order_item"
		WHERE order_id = $1 AND deleted_at = 0
	`

	err := o.db.QueryRow(query, req.Id).Scan(&res)
	if err != nil {
		return 0, errors.Wrap(err, "error while getting order items len")
	}

	return res, nil
}

func (o *orderRepo) UpsertOrderDiscount(req *order_service.UpsertOrderDiscountRequest) (*common.ResponseID, error) {

	query := `
		UPDATE
			"order"
		SET
			custom_discount_type = $2,
			custom_discount_value = $3
		WHERE
			id = $1 AND deleted_at = 0
	`

	res, err := o.db.Exec(query, req.OrderId, req.Type, req.Value)
	if err != nil {
		return nil, errors.Wrap(err, "error while update order discount")
	}

	i, err := res.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(errors.New("order not found. RowsAffected"), "error while update order discount")
	}

	if i == 0 {
		return nil, errors.New("Order not found")
	}

	return &common.ResponseID{Id: req.OrderId}, nil
}

func (o *orderRepo) UpdateOrderItemsDiscount(discount *order_service.UpsertOrderDiscountRequest, itemsLen int) error {

	var discountValue = discount.Value

	query := `
	UPDATE "order_item"
		SET
			custom_discount_type = $2,
			custom_discount_value = $3,
	`

	if discount.Type == config.OrderDiscountTypeNone {
		discountValue = 0

		query += `total_discount_price = 0`

	} else if discount.Type == config.OrderDiscountTypeAmount {

		discountValue = discountValue / float32(itemsLen)

		query += fmt.Sprintf(`total_discount_price = %v`, discountValue)

	} else if discount.Type == config.OrderDiscountTypePercentage {

		query += fmt.Sprintf(`total_discount_price = total_price*%v`, discountValue/100)

	} else {
		return errors.Wrap(errors.New("discount type must enum"), "error while chech order discount type")
	}

	query += ` WHERE order_id = $1 AND deleted_at = 0`

	_, err := o.db.Exec(query, discount.OrderId, discount.Type, discountValue)
	if err != nil {
		return errors.Wrap(err, "error while update order items discount")
	}

	return nil
}

func (o *orderRepo) CreateOrderItem(entity *order_service.CreateOrderItemRequest) (*common.ResponseID, error) {

	var (
		sellerId sql.NullString
		id       = uuid.New().String()
	)

	query := `
		INSERT INTO
			"order_item"
		(
			id,
			value,
			seller_id,
			order_id,
			product_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (order_id, product_id, deleted_at) DO
		UPDATE SET
			value = $2;
	`

	if entity.SellerId != "" {
		sellerId.Valid = true
		sellerId.String = entity.SellerId
	}

	_, err := o.db.Exec(
		query,
		id,
		entity.Value,
		sellerId,
		entity.OrderId,
		entity.ProductId,
		entity.Request.UserId,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while create order item")
	}

	return &common.ResponseID{Id: entity.OrderId}, nil
}

func (o *orderRepo) UpdateOrderItem(entity *order_service.UpdateOrderItemRequest) (*common.ResponseID, error) {

	var (
		sellerId sql.NullString
		orderId  string
	)

	query := `
	  UPDATE
			"order_item"
	  SET
			value = $2,
			seller_id = $3,
	`
	if _, err := uuid.Parse(entity.SellerId); err == nil {
		sellerId.Valid = true
		sellerId.String = entity.SellerId
	}

	if entity.Discount == nil {
		return nil, errors.Wrap(errors.New("discount type must enum"), "error while chech order discount type")
	}

	if entity.Discount.Type == config.OrderDiscountTypeNone {

		query += `
			custom_discount_type = '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'::order_discount_type,
			custom_discount_value = 0
		`

	} else if entity.Discount.Type == config.OrderDiscountTypeAmount || entity.Discount.Type == config.OrderDiscountTypePercentage {

		query += fmt.Sprintf(`
			custom_discount_type = '%s'::order_discount_type,
			custom_discount_value = %v
		`,
			entity.Discount.Type,
			entity.Discount.Value,
		)

	} else {
		return nil, errors.Wrap(errors.New("discount type must enum"), "error while chech order discount type")
	}

	query += `
		WHERE
			id = $1 AND deleted_at = 0
		RETURNING order_id
		`

	err := o.db.QueryRow(query, entity.Id, entity.Value, sellerId).Scan(&orderId)
	if err != nil {
		return nil, errors.Wrap(err, "error while update order item")
	}

	return &common.ResponseID{Id: orderId}, nil
}

func (o *orderRepo) DeleteOrderItem(req *common.RequestID) (*common.ResponseID, error) {

	var orderId string

	query := `
		UPDATE
			"order_item"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE
			id = $1 AND deleted_at = 0
		RETURNING order_id
	`

	err := o.db.QueryRow(query, req.Id).Scan(&orderId)
	if err != nil {
		return nil, errors.Wrap(err, "error while delete order item")
	}

	return &common.ResponseID{Id: orderId}, nil
}

func (o *orderRepo) CreateOrderPays(entity *order_service.CreateOrderPaysRequest) (*common.ResponseID, error) {

	var (
		values = []interface{}{}
	)

	if len(entity.Pays) < 1 {
		return nil, errors.Wrap(errors.New("pays must at least 1"), "pays is empty")
	}
	// insert pays
	query := `
			INSERT INTO
				"transaction"
			(
				id,
				value,
				order_id,
				payment_type_id,
				created_by
			)
			VALUES 
		`

	for _, pay := range entity.Pays {
		query += "(?, ?, ?, ?, ?),"
		values = append(values,
			uuid.New().String(),
			pay.Value,
			entity.OrderId,
			pay.PaymentTypeId,
			entity.Request.UserId,
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := o.db.Prepare(query)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert order pays. Prepare query")
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return nil, errors.Wrap(err, "error while insert order pays. Exec values")
	}

	stmt.Close()

	return &common.ResponseID{Id: entity.OrderId}, nil
}

func (o *orderRepo) GetOrderById(req *common.RequestID) (*order_service.GetOrderByIDResponse, error) {

	var (
		res = order_service.GetOrderByIDResponse{
			Shop:      &order_service.OrderShop{},
			Status:    &order_service.OrderStatus{},
			Client:    &order_service.OrderClient{},
			Cashbox:   &order_service.OrderCashbox{},
			Company:   &common.ShortCompany{},
			CreatedBy: &common.ShortUser{},
			Items:     make([]*order_service.GetOrderItemResponse, 0),
			Pays:      make([]*order_service.GetOrderPayResponse, 0),
			Discount:  &order_service.OrderDiscountResponse{},
		}
		payedTime         sql.NullString
		shortUser         models.NullShortUser
		client            models.NullClient
		cashbox           models.NullCashbox
		shop              models.NullShop
		company           models.NullCompany
		orderStatus       models.NullOrderStatus
		statusTranslation []byte
	)

	query := `
		SELECT
			o.id,
			o.external_id,
			o.total_price,
			o.product_sort_count,
			o.payed_time,
			o.total_discount_price,
			o.custom_discount_type,
			o.custom_discount_value,
			os.id,
			os.name,
			os.translation,
			u.id,
			u.first_name,
			u.last_name,
			c.id,
			c.first_name,
			c.last_name,
			c.phone_number,
			cb.id,
			cb.company_id,
			cb.shop_id,
			cb.title,
			sh.id,
			sh.company_id,
			sh.title,
			comp.id,
			comp.name
		FROM "order" o
		LEFT JOIN "user" u ON u.id = o.created_by AND u.deleted_at = 0
		LEFT JOIN "client" c ON c.id = o.client_id AND c.deleted_at = 0
		LEFT JOIN "order_status" os ON o.status = os.id
		LEFT JOIN "cashbox" cb ON o.cashbox_id = cb.id AND cb.deleted_at = 0
		LEFT JOIN "shop" sh ON o.shop_id = sh.id AND sh.deleted_at = 0
		LEFT JOIN "company" comp ON o.company_id = comp.id AND comp.deleted_at = 0
		WHERE
			o.id = $1 AND o.company_id = $2 AND o.deleted_at = 0
	`

	err := o.db.QueryRow(query, req.Id, req.Request.CompanyId).Scan(
		&res.Id,
		&res.ExternalId,
		&res.TotalPirce,
		&res.ProductSortCount,
		&payedTime,
		&res.Discount.Price,
		&res.Discount.Type,
		&res.Discount.Value,
		&orderStatus.Id,
		&orderStatus.Name,
		&statusTranslation,
		&shortUser.ID,
		&shortUser.FirstName,
		&shortUser.LastName,
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.PhoneNumber,
		&cashbox.Id,
		&cashbox.CompanyId,
		&cashbox.ShopId,
		&cashbox.Title,
		&shop.Id,
		&shop.CompanyId,
		&shop.Title,
		&company.Id,
		&company.Name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while get order by id")
	}

	if shortUser.ID.Valid {
		res.CreatedBy = &common.ShortUser{
			Id:        shortUser.ID.String,
			FirstName: shortUser.FirstName.String,
			LastName:  shortUser.LastName.String,
		}
	}

	if orderStatus.Id.Valid {
		res.Status = &order_service.OrderStatus{
			Id:         orderStatus.Id.String,
			Name:       orderStatus.Name.String,
			Tranlation: map[string]string{},
		}

		if err = json.Unmarshal(statusTranslation, &res.Status.Tranlation); err != nil {
			return nil, errors.Wrap(err, "error while Unmarshal order NameTrasnlation")
		}
	}

	if cashbox.Id.Valid {
		res.Cashbox = &order_service.OrderCashbox{
			Id:        cashbox.Id.String,
			Title:     cashbox.Title.String,
			ShopId:    cashbox.ShopId.String,
			CompanyId: cashbox.CompanyId.String,
		}
	}

	if client.Id.Valid {
		res.Client = &order_service.OrderClient{
			Id:          client.Id.String,
			FirstName:   client.FirstName.String,
			LastName:    client.LastName.String,
			PhoneNumber: client.PhoneNumber.String,
			CompanyId:   client.CompanyId.String,
		}
	}

	if shop.Id.Valid {
		res.Shop = &order_service.OrderShop{
			Id:        shop.Id.String,
			Title:     shop.Title.String,
			CompanyId: shop.CompanyId.String,
		}
	}

	if company.Id.Valid {
		res.Company = &common.ShortCompany{
			Id:   company.Id.String,
			Name: company.Name.String,
		}
	}

	items, err := o.getOrderItems([]string{res.Id})
	if err != nil {
		return nil, err
	}

	res.Items = items[res.Id]

	pays, err := o.getOrderPays([]string{res.Id})
	if err != nil {
		return nil, err
	}

	res.CreateTime = payedTime.String

	res.Pays = pays[res.Id]

	return &res, nil
}

func (o *orderRepo) GetAllOrders(req *order_service.GetOrdersRequest) (*order_service.GetAllOrdersResponse, error) {

	var (
		res = order_service.GetAllOrdersResponse{
			Total:      0,
			Data:       make([]*order_service.GetOrderByIDResponse, 0),
			Statistics: &order_service.OrderStatisticsResponse{},
		}
		searchFields = map[string]interface{}{
			"company_id":  req.Request.CompanyId,
			"limit":       req.Limit,
			"offset":      req.Limit * (req.Page - 1),
			"search":      req.Search,
			"status":      req.Status,
			"max_amount":  req.MaxAmount,
			"min_amount":  req.MinAmount,
			"shop_ids":    pq.Array(req.ShopIds),
			"seller_ids":  pq.Array(req.SellerIds),
			"cashier_ids": pq.Array(req.CashierIds),
			"client_ids":  pq.Array(req.ClientIds),
			"start_date":  req.StartDate,
			"end_date":    req.EndDate,
		}
		orderIds = make([]string, 0)
	)

	namedQuery := `
		SELECT
			o.id,
			o.external_id,
			o.total_price,
			o.total_discount_price,
			o.product_sort_count,
			o.payed_time,
			o.cashier_id,
			os.id,
			os.name,
			os.translation,
			u.id,
			u.first_name,
			u.last_name,
			c.id,
			c.first_name,
			c.last_name,
			c.phone_number,
			cb.id,
			cb.company_id,
			cb.shop_id,
			cb.title,
			sh.id,
			sh.company_id,
			sh.title,
			comp.id,
			comp.name
		FROM
			"order" o
		LEFT JOIN "user" u ON u.id = o.created_by AND u.deleted_at = 0
		LEFT JOIN "order_status" os ON o.status = os.id
		LEFT JOIN "client" c ON c.id = o.client_id AND c.deleted_at = 0
		LEFT JOIN "cashbox" cb ON o.cashbox_id = cb.id AND cb.deleted_at = 0
		LEFT JOIN "shop" sh ON o.shop_id = sh.id AND sh.deleted_at = 0
		LEFT JOIN "company" comp ON o.company_id = comp.id AND comp.deleted_at = 0
	`

	filter := `
		WHERE
			o.company_id = :company_id AND
			o.deleted_at = 0 AND
			o.status != '7069e210-7d2e-4a12-9160-3ef82f18ef4d'
	`

	if req.Search != "" {
		filter += `
		AND o.external_id ILIKE '%' || :search || '%'
	`
	}

	if req.Status != "" {
		filter += ` AND o.status = :status `
	}

	if req.MaxAmount != 0 {
		filter += ` AND o.total_price <= :max_amount `
	}

	if req.MinAmount != 0 {
		filter += ` AND o.total_price >= :min_amount `
	}

	if req.StartDate != "" {
		filter += ` AND (
			o."payed_time" >= :start_date 
		)
		`
	}

	if req.EndDate != "" {
		filter += ` AND (
			o."payed_time" <= :end_date 
		)
		`
	}

	if len(req.ShopIds) > 0 {
		filter += ` AND o.shop_id = ANY(:shop_ids) `
	}

	if len(req.SellerIds) > 0 {
		filter += ` AND o.seller_id = ANY(:seller_ids) `
	}

	if len(req.CashierIds) > 0 {
		filter += ` AND o.cashier_id = ANY(:cashier_ids)`
	}

	if len(req.ClientIds) > 0 {
		filter += ` AND o.client_id=ANY(:client_ids) `
	}

	namedQuery += filter + `
		ORDER BY o.created_at DESC
		LIMIT :limit
		OFFSET :offset
	`

	rows, err := o.db.NamedQuery(namedQuery, searchFields)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting orders. NamedQuery")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			order = order_service.GetOrderByIDResponse{
				Shop:      &order_service.OrderShop{},
				Status:    &order_service.OrderStatus{},
				Client:    &order_service.OrderClient{},
				Cashbox:   &order_service.OrderCashbox{},
				Company:   &common.ShortCompany{},
				CreatedBy: &common.ShortUser{},
				Items:     make([]*order_service.GetOrderItemResponse, 0),
				Pays:      make([]*order_service.GetOrderPayResponse, 0),
				Discount:  &order_service.OrderDiscountResponse{},
			}
			shortUser         models.NullShortUser
			client            models.NullClient
			payedTime         sql.NullString
			cashbox           models.NullCashbox
			shop              models.NullShop
			company           models.NullCompany
			orderStatus       models.NullOrderStatus
			statusTranslation []byte
		)

		err = rows.Scan(
			&order.Id,
			&order.ExternalId,
			&order.TotalPirce,
			&order.Discount.Price,
			&order.ProductSortCount,
			&payedTime,
			&order.CashierId,
			&orderStatus.Id,
			&orderStatus.Name,
			&statusTranslation,
			&shortUser.ID,
			&shortUser.FirstName,
			&shortUser.LastName,
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.PhoneNumber,
			&cashbox.Id,
			&cashbox.CompanyId,
			&cashbox.ShopId,
			&cashbox.Title,
			&shop.Id,
			&shop.CompanyId,
			&shop.Title,
			&company.Id,
			&company.Name,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning rows")
		}

		if shortUser.ID.Valid {
			order.CreatedBy = &common.ShortUser{
				Id:        shortUser.ID.String,
				FirstName: shortUser.FirstName.String,
				LastName:  shortUser.LastName.String,
			}
		}

		if orderStatus.Id.Valid {
			order.Status = &order_service.OrderStatus{
				Id:         orderStatus.Id.String,
				Name:       orderStatus.Name.String,
				Tranlation: map[string]string{},
			}

			if err = json.Unmarshal(statusTranslation, &order.Status.Tranlation); err != nil {
				return nil, errors.Wrap(err, "error while Unmarshal order NameTrasnlation")
			}
		}

		if cashbox.Id.Valid {
			order.Cashbox = &order_service.OrderCashbox{
				Id:        cashbox.Id.String,
				Title:     cashbox.Title.String,
				ShopId:    cashbox.ShopId.String,
				CompanyId: cashbox.CompanyId.String,
			}
		}

		if client.Id.Valid {
			order.Client = &order_service.OrderClient{
				Id:          client.Id.String,
				FirstName:   client.FirstName.String,
				LastName:    client.LastName.String,
				PhoneNumber: client.PhoneNumber.String,
				CompanyId:   client.CompanyId.String,
			}
		}

		if shop.Id.Valid {
			order.Shop = &order_service.OrderShop{
				Id:        shop.Id.String,
				Title:     shop.Title.String,
				CompanyId: shop.CompanyId.String,
			}
		}

		if company.Id.Valid {
			order.Company = &common.ShortCompany{
				Id:   company.Id.String,
				Name: company.Name.String,
			}
		}

		order.CreateTime = payedTime.String

		orderIds = append(orderIds, order.Id)
		res.Data = append(res.Data, &order)
	}

	countQuery := `
		SELECT 
			count(*) AS total
		FROM 
			"order" o
	` +
		filter

	resStmt, err := o.db.PrepareNamed(countQuery)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning total PrepareNamed")
	}

	defer resStmt.Close()

	if err != resStmt.QueryRow(searchFields).Scan(&res.Total) {
		return nil, errors.Wrap(err, "error while scanning total")
	}

	items, err := o.getOrderItems(orderIds)
	if err != nil {
		return nil, err
	}

	pays, err := o.getOrderPays(orderIds)
	if err != nil {
		return nil, err
	}

	for _, order := range res.Data {
		order.Items = items[order.Id]
		order.Pays = pays[order.Id]
	}

	statisticsQuery := `
		SELECT COALESCE(SUM(oi.total_price), 0) AS total_amount, COALESCE(COUNT(*),0) as total_count
		FROM "order" o
		LEFT JOIN order_item oi on oi.order_id = o.id AND oi.deleted_at = 0
	`
	statisticsQuery += filter

	resStmt, err = o.db.PrepareNamed(statisticsQuery)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning order statistics PrepareNamed")
	}

	defer resStmt.Close()

	if err := resStmt.QueryRow(searchFields).Scan(&res.Statistics.TransactionsAmount, &res.Statistics.TransactionsCount); err != nil {
		return nil, errors.Wrap(err, "error while scanning order statistics")
	}

	return &res, nil
}

func (o *orderRepo) DeleteOrderById(req *common.RequestID) (*common.ResponseID, error) {

	query := `
		UPDATE
			"order"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE
			id = $1 AND deleted_at = 0
	`

	res, err := o.db.Exec(
		query,
		req.Id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while delete order")
	}

	i, err := res.RowsAffected()
	if err != nil {
		return nil, errors.New("order item not found. RowsAffected")
	}

	if i == 0 {
		return nil, errors.New("order not found")
	}

	return &common.ResponseID{Id: req.Id}, nil
}

func (o *orderRepo) getOrderItems(ids []string) (map[string][]*order_service.GetOrderItemResponse, error) {

	var (
		res     = make(map[string][]*order_service.GetOrderItemResponse)
		product models.NullProduct
	)

	query := `
		SELECT
			o_i.id,
			o_i.price,
			o_i.value,
			o_i.total_price,
			o_i.order_id,
			o_i.product_id,
			o_i.total_discount_price,
			o_i.custom_discount_type,
			o_i.custom_discount_value,
			seller.id,
			seller.first_name,
			seller.last_name,
			p.name,
			p.sku,
			p.mxik_code,
			p.image,
			mu.id,
			mu.short_name,
			mu.long_name,
			u.id,
			u.first_name,
			u.last_name
		FROM
			"order_item" o_i
		LEFT JOIN "user" u ON u.id = o_i.created_by AND u.deleted_at = 0
		LEFT JOIN "user" seller ON seller.id = o_i.seller_id AND seller.deleted_at = 0
		LEFT JOIN "product" p ON p.id = o_i.product_id AND p.deleted_at = 0
		LEFT JOIN "measurement_unit" mu ON mu.id = p.measurement_unit_id AND mu.deleted_at = 0
		WHERE
			o_i.order_id = ANY ($1) AND o_i.deleted_at = 0
		ORDER BY o_i.created_at DESC
	`

	rows, err := o.db.Query(query, pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "error while getting order items")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			shortUser models.NullShortUser
			seller    models.NullShortUser
			orderItem = order_service.GetOrderItemResponse{
				Discount: &order_service.OrderDiscountResponse{},
			}
			measurementUnit models.NullMeasurementUnit
		)

		err = rows.Scan(
			&orderItem.Id,
			&orderItem.Price,
			&orderItem.Value,
			&orderItem.TotalPrice,
			&orderItem.OrderId,
			&orderItem.ProductId,
			&orderItem.Discount.Price,
			&orderItem.Discount.Type,
			&orderItem.Discount.Value,
			&seller.ID,
			&seller.FirstName,
			&seller.LastName,
			&product.Name,
			&product.Sku,
			&product.MxikCode,
			&product.Image,
			&measurementUnit.Id,
			&measurementUnit.ShortName,
			&measurementUnit.LongName,
			&shortUser.ID,
			&shortUser.FirstName,
			&shortUser.LastName,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting categories childs rows.Scan")
		}

		orderItem.NewPrice = (orderItem.TotalPrice - orderItem.Discount.Price) / orderItem.Value

		orderItem.MeasurementUnit = &catalog_service.ShortMeasurementUnit{
			Id:        measurementUnit.Id.String,
			ShortName: measurementUnit.ShortName.String,
			LongName:  measurementUnit.LongName.String,
		}

		orderItem.CreatedBy = &common.ShortUser{
			Id:        shortUser.ID.String,
			FirstName: shortUser.FirstName.String,
			LastName:  shortUser.LastName.String,
		}

		orderItem.Seller = &common.ShortUser{
			Id:        seller.ID.String,
			FirstName: seller.FirstName.String,
			LastName:  seller.LastName.String,
		}

		orderItem.ProductName = product.Name.String
		orderItem.Image = product.Image.String
		orderItem.Sku = product.Sku.String
		orderItem.MxikCode = product.MxikCode.String

		res[orderItem.OrderId] = append(res[orderItem.OrderId], &orderItem)
	}

	return res, nil
}

func (o *orderRepo) getOrderPays(orderIds []string) (map[string][]*order_service.GetOrderPayResponse, error) {

	var (
		res = make(map[string][]*order_service.GetOrderPayResponse)
	)

	query := `
		SELECT
			tr.id,
			tr.value,
			tr.order_id,
			pt.id,
			pt.company_id,
			pt.name,
			u.id,
			u.first_name,
			u.last_name
		FROM
			"transaction" tr
		LEFT JOIN "user" u ON u.id = tr.created_by AND u.deleted_at = 0
		LEFT JOIN "payment_type" pt ON pt.id = tr.payment_type_id AND pt.deleted_at = 0
		WHERE
			tr.deleted_at = 0 AND tr.order_id = ANY ($1)
	`

	rows, err := o.db.Query(query, pq.Array(orderIds))
	if err != nil {
		return nil, errors.Wrap(err, "error while getting order pays")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			shortUser   models.NullShortUser
			orderPay    order_service.GetOrderPayResponse
			paymentType models.NullPaymentType
		)

		err = rows.Scan(
			&orderPay.Id,
			&orderPay.Value,
			&orderPay.OrderId,
			&paymentType.Id,
			&paymentType.Name,
			&paymentType.CompanyId,
			&shortUser.ID,
			&shortUser.FirstName,
			&shortUser.LastName,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting categories childs rows.Scan")
		}

		if shortUser.ID.Valid {
			orderPay.CreatedBy = &common.ShortUser{
				Id:        shortUser.ID.String,
				FirstName: shortUser.FirstName.String,
				LastName:  shortUser.LastName.String,
			}
		}

		if paymentType.Id.Valid {
			orderPay.PaymentType = &order_service.TransactionPaymentType{
				Id:        paymentType.Id.String,
				Name:      paymentType.Name.String,
				CompanyId: paymentType.CompanyId.String,
			}
		}

		res[orderPay.OrderId] = append(res[orderPay.OrderId], &orderPay)
	}

	return res, nil
}

func (o *orderRepo) GetPaymentTypeAnalytics(req *report_service.GetDashboardAnalyticsReq) (res []*report_service.Payments, err error) {
	var (
		filter = ""
		args   = make(map[string]interface{})
	)
	queryShops := `
	SELECT 
		sh."title",
		o."shop_id"
	FROM "payment_type" AS p
	JOIN "transaction" AS t ON p."id" = t."payment_type_id"
	JOIN "order" AS o ON t."order_id" = o."id"
	JOIN "shop" AS sh ON o."shop_id" = sh."id"
	WHERE o."company_id" = $1
	GROUP BY 
		o."shop_id",
		sh."title"
	`

	rows, err := o.db.Query(queryShops, req.Request.CompanyId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return nil, errors.Wrap(err, "error while getting shops for GetPaymentTypeAnalytics")
		}
	}
	defer rows.Close()
	for rows.Next() {
		var (
			shop     = common.ShortShop{}
			payments = report_service.Payments{}
		)
		err = rows.Scan(
			&shop.Name,
			&shop.Id,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting shops for GetPaymentTypeAnalytics rows.Scan")
		}
		payments.Shop = &shop
		payments.TotalPrice = 0
		res = append(res, &payments)
	}

	if len(req.DateFrom) > 0 {
		if dateFrom, err := helper.ParseDate(req.DateFrom); err == nil {
			args["dateFrom"] = dateFrom
			filter += ` AND o."payed_time" >= :dateFrom`
		} else if err != nil {
			return nil, err
		}
	}
	if len(req.DateTo) > 0 {
		if dateTo, err := helper.ParseDate(req.DateTo); err == nil {
			args["dateTo"] = dateTo
			filter += ` AND o."payed_time" <= :dateTo`
		} else if err != nil {
			return nil, err
		}
	}

	args["company_id"] = req.Request.CompanyId

	for _, v := range res {
		args["shop_id"] = v.Shop.Id
		query := `
		SELECT 
			p."name",
			sum(t."value")
		FROM "payment_type" AS p
		JOIN "transaction" AS t ON p."id" = t."payment_type_id"
		JOIN "order" AS o ON t."order_id" = o."id"
		WHERE	o."company_id" = :company_id
				AND o."shop_id" = :shop_id
				AND o."status" = 'd3bde6a2-532c-4f08-811f-0385e804c885'
		`
		query += filter
		query += `
		GROUP BY p."name"`

		query, arrArgs := helper.ReplaceQueryParams(query, args)
		rows, err = o.db.Query(query, arrArgs...)
		if err != nil {
			if err.Error() == "no rows in result set" {
				err = nil
			} else {
				return nil, errors.Wrap(err, "error while getting GetPaymentTypeAnalytics")
			}
		}
		defer rows.Close()

		for rows.Next() {
			var pt report_service.PaymentsType
			err = rows.Scan(
				&pt.Name,
				&pt.Price,
			)
			if err != nil {
				return nil, errors.Wrap(err, "error while getting GetPaymentTypeAnalytics rows.Scan")
			}
			v.TotalPrice += pt.Price
			v.PaymentsTypes = append(v.PaymentsTypes, &pt)
		}
	}
	return res, nil
}

func (o *orderRepo) AddClientToOrder(in *order_service.AddClientToOrderRequest) (*order_service.GetOrderByIDResponse, error) {

	query := `
		UPDATE "order" SET "client_id"=$1 WHERE id=$2 AND company_id=$3 AND status=$4 AND deleted_at=0
	`

	res, err := o.db.Exec(query, in.ClientId, in.OrderId, in.Request.CompanyId, config.OrderStatusDraft)
	if err != nil {
		return nil, err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	order, err := o.GetOrderById(&common.RequestID{Request: in.Request, Id: in.OrderId})
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderRepo) RemoveClientFromOrder(in *order_service.RemoveClientFromOrderRequest) (*order_service.GetOrderByIDResponse, error) {
	query := `
	UPDATE "order" SET "client_id"=NULL WHERE id=$1 AND company_id=$2 AND status=$3 AND deleted_at=0
`

	res, err := o.db.Exec(query, in.OrderId, in.Request.CompanyId, config.OrderStatusDraft)
	if err != nil {
		return nil, err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	order, err := o.GetOrderById(&common.RequestID{Request: in.Request, Id: in.OrderId})
	if err != nil {
		return nil, err
	}

	return order, nil
}
