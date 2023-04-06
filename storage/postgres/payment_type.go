package postgres

import (
	"genproto/common"
	"strings"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/helper"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type paymentTypeRepo struct {
	db  models.DB
	log logger.Logger
}

func NewPaymentTypeRepo(db models.DB, log logger.Logger) repo.PaymentTypeI {
	return &paymentTypeRepo{
		db:  db,
		log: log,
	}
}

func (c *paymentTypeRepo) Upsert(entity *common.CommonPaymentTypes) error {
	var query = `
		INSERT INTO
				"payment" 
			(
					id,
					name
			)
			VALUES 
			(
				$1,
				$2

			);
		`
	entity.Id = uuid.New().String()

	_, err := c.db.Query(
		query,
		entity.Id,
		entity.Name,
	)
	if err != nil {
		return errors.Wrap(err, "error while inserting payment type")
	}
	return err
}

func (c *paymentTypeRepo) UpsertMany(req []*common.CommonPaymentTypes) error {
	
	var values = []interface{}{}

	query := `
		INSERT INTO "payment_type"
		(id, name, company_id, created_by)
		VALUES
	`

	for _, payment_types := range req {
		query += "(?, ?, ?, ?),"
		values = append(values,
			payment_types.Id,
			payment_types.Name,
			payment_types.Request.CompanyId,
			payment_types.Request.UserId,
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while upsert payment_types. Prepare")
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while upsert cheques. Exec")
	}

	stmt.Close()

	return nil
}
