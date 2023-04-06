package postgres

import (
	"encoding/json"
	"genproto/common"
	"genproto/corporate_service"
	"strings"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/helper"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type chequeRepo struct {
	db  models.DB
	log logger.Logger
}

func NewChequeRepo(db models.DB, log logger.Logger) repo.ChequeI {
	return &chequeRepo{
		db:  db,
		log: log,
	}
}

func UpsertFields(db models.DB, req []*common.ChequeCopyRequest) error {

	var values = []interface{}{}

	if len(req) <= 0 {
		return errors.New("fields must at least one")
	}

	query := `
		INSERT INTO "cheque_field"
		("field_id", "cheque_id", "position")
		VALUES 
	`

	for _, cheque := range req {

		for _, fieldId := range cheque.ChequeFields {

			query += "(?, ?, ?),"

			values = append(values, fieldId.FieldId, cheque.Id, fieldId.Position)
		}
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	query += ` ON CONFLICT (field_id, cheque_id, deleted_at) DO UPDATE SET position=EXCLUDED.position`

	stmt, err := db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheque_fields. Prepare")
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheque_fields. Exec")
	}

	return nil
}

func UpsertLogo(db models.DB, req *common.ChequeLogoCopyRequest) error {

	query := `
		INSERT INTO "cheque_logo"
		("image", "cheque_id", "left", "right", "top", "bottom")
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (cheque_id) DO
		UPDATE
			SET
			"image" = $1,
			"left" = $3,
			"right" = $4,
			"top" = $5,
			"bottom" = $6;
	`

	_, err := db.Exec(query, req.Image, req.ChequeId, req.Left, req.Right, req.Top, req.Bottom)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheque logo")
	}

	return nil
}

func (c *chequeRepo) Upsert(entity *common.ChequeCopyRequest) error {

	query := `
		INSERT INTO "cheque"
		(id, company_id, name, message)
		VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO UPDATE SET name=$3, message=$4
	`

	_, err := c.db.Exec(
		query,
		entity.Id,
		entity.CompanyId,
		entity.Name,
		entity.Message,
	)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheque")
	}

	if err := UpsertFields(c.db, []*common.ChequeCopyRequest{entity}); err != nil {
		return errors.Wrap(err, "error while upsert cheque fields")
	}

	if err := UpsertLogo(c.db, entity.ChequeLogo); err != nil {
		return errors.Wrap(err, "error while upsert logo")
	}

	return nil
}

func (c *chequeRepo) UpsertMany(req []*common.ChequeCopyRequest) error {

	var values = []interface{}{}

	query := `
		INSERT INTO "cheque"
		(id, company_id, name, message)
		VALUES
	`

	for _, measurement_unit := range req {
		query += "(?, ?, ?, ?),"
		values = append(values,
			measurement_unit.Id,
			measurement_unit.CompanyId,
			measurement_unit.Name,
			measurement_unit.Message,
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheques. Prepare")
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheques. Exec")
	}

	stmt.Close()

	if err := UpsertFields(c.db, req); err != nil {
		return err
	}

	return nil
}

func (c *chequeRepo) GetById(entity *common.RequestID) (*models.GetChequeResponse, error) {

	var (
		res       models.GetChequeResponse
		blocksMap map[string]*models.GetReceiptBlockResponse = make(map[string]*models.GetReceiptBlockResponse)
		blockIds                                             = make([]string, 0)
		logo      models.NullChequeLogo
	)

	query := `
		SELECT 
			c."id",
			c."name",
			c."message",
			chl."image",
			chl."top",
			chl."right",
			chl."bottom",
			chl."left"
		FROM "cheque" c
		LEFT JOIN "cheque_logo" chl ON c."id" = chl."cheque_id"
		WHERE c."id" = $1 AND c."deleted_at" = 0 AND c."company_id" = $2;
	`

	err := c.db.QueryRow(query, entity.Id, entity.Request.CompanyId).Scan(
		&res.Id,
		&res.Name,
		&res.Message,
		&logo.Image,
		&logo.Top,
		&logo.Right,
		&logo.Bottom,
		&logo.Left,
	)
	if err != nil {
		return nil, errors.New("error while getting cheque QueryRow in GetById")
	}

	if logo.Image.Valid {
		res.Logo = &models.ChequeLogo{
			Image:  logo.Image.String,
			Left:   int8(logo.Left.Int16),
			Right:  int8(logo.Right.Int16),
			Top:    int8(logo.Top.Int16),
			Bottom: int8(logo.Bottom.Int16),
		}
	}

	query = `
		SELECT 
			"id",
			"name",
			"name_tr"
		FROM "receipt_block" rb
		WHERE "deleted_at" = 0;
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, errors.New("error while getting block Query")
	}

	defer rows.Close()

	for rows.Next() {

		block := models.GetReceiptBlockResponse{
			Fields: make([]*models.GetReceiptFieldResponse, 0),
		}

		nameTranslation := []byte{}

		if err := rows.Scan(&block.Id, &block.Name, &nameTranslation); err != nil {
			return nil, errors.New("error while getting block Scan")
		}

		if len(nameTranslation) > 0 {
			if err := json.Unmarshal(nameTranslation, &block.NameTr); err != nil {
				return nil, errors.New("error while unmarshaling block translate")
			}
		}

		blocksMap[block.Id] = &block
		blockIds = append(blockIds, block.Id)
	}

	query = `
		SELECT 
			rf."id",
			rf."name",
			rf."name_tr",
			rf."block_id"
		FROM receipt_field rf
		WHERE "deleted_at" = 0 AND "id" = ANY($2);
	`

	rows, err = c.db.Query(query, entity.Id, pq.Array(blockIds))
	if err != nil {
		return nil, errors.New("error while getting receipt_field Query")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			field           = models.GetReceiptFieldResponse{}
			blockID         string
			nameTranslation = []byte{}
		)

		if err := rows.Scan(&field.Id, &field.Name, &nameTranslation, &blockID); err != nil {
			return nil, errors.New("error while Scan field")
		}

		block, ok := blocksMap[blockID]
		if ok {
			block.Fields = append(block.Fields, &field)
			blocksMap[blockID] = block
		}
	}

	for _, block := range blocksMap {
		res.Blocks = append(res.Blocks, block)
	}

	return &res, nil
}

func (c *chequeRepo) GetByOrderId(orderId string) (*corporate_service.Cheque, error) {

	var (
		res       corporate_service.Cheque
		blocksMap = make(map[string]*corporate_service.RecieptBlock)
		blockIds  = make([]string, 0)
		logo      models.NullChequeLogo
	)

	query := `
		SELECT 
			c."id",
			c."name",
			c."message",
			chl."image",
			chl."top",
			chl."right",
			chl."bottom",
			chl."left"
		FROM "cheque" c
		JOIN "order" o ON o."id" = $1 
		JOIN "cashbox" cashb ON cashb."id" = o."cashbox_id"
		LEFT JOIN "cheque_logo" chl ON chl."cheque_id" = c."id" 
		WHERE c."id" = cashb."cheque_id";
	`

	err := c.db.QueryRow(query, orderId).Scan(
		&res.Id,
		&res.Name,
		&res.Message,
		&logo.Image,
		&logo.Top,
		&logo.Right,
		&logo.Bottom,
		&logo.Left,
	)
	if err != nil {
		return nil, errors.Wrap(err, "order cheque in GetByOrderId")
	}

	if logo.Image.Valid {
		res.Logo = &corporate_service.ChequeLogo{
			Image:  logo.Image.String,
			Left:   float32(logo.Left.Int16),
			Right:  float32(logo.Right.Int16),
			Top:    float32(logo.Top.Int16),
			Bottom: float32(logo.Bottom.Int16),
		}
	}

	query = `
		SELECT 
			"id",
			"name",
			"name_tr"
		FROM "receipt_block" rb
		WHERE "deleted_at" = 0;
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting blocks Query")
	}

	defer rows.Close()

	for rows.Next() {

		block := corporate_service.RecieptBlock{
			Fields: make([]*corporate_service.RecieptField, 0),
		}

		nameTranslation := []byte{}

		if err := rows.Scan(&block.Id, &block.Name, &nameTranslation); err != nil {
			return nil, errors.Wrap(err, "error while getting blocks Scan")
		}

		if len(nameTranslation) > 0 {
			if err := json.Unmarshal(nameTranslation, &block.NameTranslation); err != nil {
				return nil, errors.New("error while unmarshaling block translation")
			}
		}

		blocksMap[block.Id] = &block
		blockIds = append(blockIds, block.Id)
	}

	query = `
		SELECT 
			rf."id",
			rf."name",
			rf."name_tr",
			rf."block_id"
		FROM receipt_field rf
		WHERE "deleted_at" = 0 AND "id" = ANY($1)
	`

	rows, err = c.db.Query(query, pq.Array(blockIds))
	if err != nil {
		return nil, errors.Wrap(err, "error while getting receipt_field Query")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			field           corporate_service.RecieptField
			blockID         string
			nameTranslation = []byte{}
		)

		if err := rows.Scan(&field.Id, &field.Name, &nameTranslation, &blockID, &field); err != nil {
			return nil, errors.New("error while scanning receipt_field Scan")
		}

		block, ok := blocksMap[blockID]
		if ok {
			block.Fields = append(block.Fields, &field)
			blocksMap[blockID] = block
		}
	}

	for _, block := range blocksMap {
		res.Blocks = append(res.Blocks, block)
	}

	return &res, nil
}
