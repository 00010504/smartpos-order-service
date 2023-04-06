package postgres

import (
	"database/sql"
	"genproto/catalog_service"
	"genproto/common"
	"strings"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/helper"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type productRepo struct {
	db  models.DB
	log logger.Logger
}

func NewProductRepo(log logger.Logger, db models.DB) repo.ProductI {
	return &productRepo{
		db:  db,
		log: log,
	}
}

func (p *productRepo) Upsert(entity *common.CreateProductCopyRequest) error {

	p.log.Info("entity", logger.Any("entity", entity))
	var values = []interface{}{}

	query := `
		INSERT INTO
			"product"
		(
			id,
			name,
			is_marking,
			sku,
			image,
			mxik_code,
			company_id,
			measurement_unit_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO
		UPDATE
		SET
			name = $2,
			is_marking = $3,
			sku = $4,
			image = $5,
			measurement_unit_id = $8,
			mxik_code = $6;
	`

	_, err := p.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.IsMarking,
		entity.Sku,
		entity.Image,
		entity.MxikCode,
		entity.Request.CompanyId,
		entity.MeasurementUnitId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert product copy")
	}

	// insert measurement values
	if len(entity.ShopMeasurementValues) > 0 {

		values = []interface{}{}

		query = `
			INSERT INTO
				"measurement_values"
			(
				shop_id,
				product_id,
				is_available,
				in_stock
			)
			VALUES 
		`

		for _, value := range entity.ShopMeasurementValues {
			query += "(?, ?, ?, ?),"
			values = append(values,
				value.ShopId,
				entity.Id,
				value.IsAvailable,
				value.InStock,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")
		query += `
		ON CONFLICT (product_id, shop_id) DO
		UPDATE
			SET
			is_available = excluded.is_available,
			in_stock = excluded.in_stock;
		`

		stmt, err := p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product measurement_values")
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			stmt.Close()
			return errors.Wrap(err, "error while insert product measurement_values")
		}

		stmt.Close()

		values = []interface{}{}

		query = `
			INSERT INTO
				"shop_price"
			(
				id,
				supply_price,
				min_price,
				max_price,
				retail_price,
				whole_sale_price,
				shop_id,
				product_id
			)
			VALUES 
		`
		for _, value := range entity.ShopMeasurementValues {
			query += "(?, ?, ?, ?, ?, ?, ?, ?),"
			values = append(values,
				uuid.New().String(),
				value.SupplyPrice,
				value.MinPrice,
				value.MaxPrice,
				value.RetailPrice,
				value.WholeSalePrice,
				value.ShopId,
				entity.Id,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += `
		ON CONFLICT (product_id, shop_id) DO
		UPDATE
			SET
			supply_price = excluded.supply_price,
			min_price = excluded.min_price,
			max_price = excluded.max_price,
			retail_price = excluded.retail_price,
			whole_sale_price = excluded.whole_sale_price;
		`

		stmt, err = p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product shop_price. Prepare")
		}
		_, err = stmt.Exec(values...)
		if err != nil {
			return errors.Wrap(err, "error while insert product shop_price. Exec")
		}

		stmt.Close()
	}

	// insert barcodes
	if len(entity.Barcode) > 0 {

		values = []interface{}{}

		query = `
			INSERT INTO
				"product_barcode"
			(
				barcode,
				product_id
			)
			VALUES 
		`

		for _, barcode := range entity.Barcode {
			query += "(?, ?),"
			values = append(values,
				barcode,
				entity.Id,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += `
		ON CONFLICT (barcode, product_id) DO NOTHING
		`

		stmt, err := p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product barcodes. Prepare")
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			stmt.Close()
			return errors.Wrap(err, "error while insert product barcodes. Exec")
		}

		stmt.Close()
	}

	return nil
}

func (p *productRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}

func (p *productRepo) UpsertShopMeasurmentValue(req *catalog_service.UpsertShopMeasurmentValueRequest) error {

	var (
		values []interface{}
	)

	query := `
			INSERT INTO
				"measurement_values"
			(
				shop_id,
				in_stock,
				product_id
			)
			VALUES
			
		`

	for _, v := range req.ProductsValues {

		query += `(?, ?, ?),`

		values = append(values, req.ShopId, v.Amount, v.ProductId)

	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	query += `
		ON CONFLICT (shop_id, product_id) DO UPDATE SET in_stock = EXCLUDED.in_stock
	`
	_, err := p.db.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "error while upsertShopMeasurementValues")
	}

	return nil
}

func (p *productRepo) InsertMany(products []*common.CreateProductCopyRequest) error {

	var (
		values            = []interface{}{}
		productBarcodes   = []interface{}{}
		measurementValues = []interface{}{}
		shopPrices        = []interface{}{}
	)

	if len(products) <= 0 {
		return nil
	}

	query := `
		INSERT INTO
			"product"
			(
				id,
				name,
				is_marking,
				sku,
				company_id,
				measurement_unit_id
			)
		VALUES
	`

	queryMeasurementValues := `
			INSERT INTO
				"measurement_values"
			(
				shop_id,
				product_id,
				is_available,
				in_stock
			)
			VALUES 
		`

	queryShopPrice := `
		INSERT INTO
			"shop_price"
		(
			id,
			supply_price,
			min_price,
			max_price,
			retail_price,
			whole_sale_price,
			shop_id,
			product_id
		)
		VALUES 
	`

	queryBarcode := `
			INSERT INTO
				"product_barcode"
			(
				barcode,
				product_id
			)
			VALUES 
		`

	for _, product := range products {
		query += "(?, ?, ?, ?, ?, ?),"
		values = append(values,
			product.Id,
			product.Name,
			product.IsMarking,
			product.Sku,
			product.Request.CompanyId,
			product.MeasurementUnitId,
		)
		for _, sshopMeasurementValues := range product.ShopMeasurementValues {
			queryMeasurementValues += "(?, ?, ?, ?),"
			measurementValues = append(measurementValues,
				sshopMeasurementValues.ShopId,
				product.Id,
				sshopMeasurementValues.IsAvailable,
				sshopMeasurementValues.InStock,
			)
		}
		for _, value := range product.ShopMeasurementValues {
			queryShopPrice += "(?, ?, ?, ?, ?, ?, ?, ?),"
			shopPrices = append(shopPrices,
				uuid.New().String(),
				value.SupplyPrice,
				value.MinPrice,
				value.MaxPrice,
				value.RetailPrice,
				value.WholeSalePrice,
				value.ShopId,
				product.Id,
			)
		}
		for _, barcode := range product.Barcode {
			queryBarcode += "(?, ?),"
			productBarcodes = append(productBarcodes,
				barcode,
				product.Id,
			)
		}
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	query += `
		ON CONFLICT (id) DO NOTHING
	`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while insertMany products. Prepare")
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while insertMany products. Exec")
	}

	if len(measurementValues) > 0 {
		queryMeasurementValues = strings.TrimSuffix(queryMeasurementValues, ",")
		queryMeasurementValues = helper.ReplaceSQL(queryMeasurementValues, "?")
		queryMeasurementValues += `
			ON CONFLICT (product_id, shop_id) DO
			UPDATE
				SET
				is_available = excluded.is_available,
				in_stock = excluded.in_stock;
			`
		stmt1, err := p.db.Prepare(queryMeasurementValues)
		if err != nil {
			return errors.Wrap(err, "error while insert many product measurement_values. PREPARE")
		}
		defer stmt1.Close()

		_, err = stmt1.Exec(measurementValues...)
		if err != nil {
			return errors.Wrap(err, "error while insert many product measurement_values. EXEC")
		}
	}

	if len(shopPrices) > 0 {
		queryShopPrice = strings.TrimSuffix(queryShopPrice, ",")
		queryShopPrice = helper.ReplaceSQL(queryShopPrice, "?")

		queryShopPrice += `
		ON CONFLICT (product_id, shop_id) DO
		UPDATE
			SET
			supply_price = excluded.supply_price,
			min_price = excluded.min_price,
			max_price = excluded.max_price,
			retail_price = excluded.retail_price,
			whole_sale_price = excluded.whole_sale_price;
		`

		stmt2, err := p.db.Prepare(queryShopPrice)
		if err != nil {
			return errors.Wrap(err, "error while insert many product shop_price. Prepare")
		}
		defer stmt2.Close()

		_, err = stmt2.Exec(shopPrices...)
		if err != nil {
			return errors.Wrap(err, "error while insert many product shop_price. Exec")
		}
	}

	if len(productBarcodes) > 0 {
		queryBarcode = strings.TrimSuffix(queryBarcode, ",")
		queryBarcode = helper.ReplaceSQL(queryBarcode, "?")

		queryBarcode += `
		ON CONFLICT (barcode, product_id) DO NOTHING
		`

		stmt3, err := p.db.Prepare(queryBarcode)
		if err != nil {
			return errors.Wrap(err, "error while insert many product barcodes. Prepare")
		}
		defer stmt3.Close()

		_, err = stmt3.Exec(productBarcodes...)
		if err != nil {
			stmt.Close()
			return errors.Wrap(err, "error while insert many product barcodes. Exec")
		}
	}
	return nil
}

func (p *productRepo) BulkUpdate(req *catalog_service.ProductBulkOperationRequest) (*common.ResponseID, error) {
	var (
		resposeID = uuid.NewString()
	)

	if req.ProductField == "name" {

		nameQuery := `
			UPDATE
				"product"
			SET
				name = $2
			WHERE
				id = ANY($1) AND deleted_at = 0
	`

		res, err := p.db.Exec(nameQuery, pq.Array(req.ProductIds), req.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error while update product name")
		}

		if i, _ := res.RowsAffected(); i == 0 {
			return nil, sql.ErrNoRows
		}
	}

	if req.ProductField == "measurement_value" {

		measurementValueQuery := `
			UPDATE
				"product" 
			SET
				measurement_unit_id = $2
			WHERE
				id = ANY($1) AND deleted_at = 0
	`

		res, err := p.db.Exec(measurementValueQuery, pq.Array(req.ProductIds), req.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error while update product measurementValue")
		}

		if i, _ := res.RowsAffected(); i == 0 {
			return nil, sql.ErrNoRows
		}
	}

	if req.ProductField == "category" {
	}

	if req.ProductField == "low_stock" {
	}

	return &common.ResponseID{Id: resposeID}, nil
}
