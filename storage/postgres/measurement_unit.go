package postgres

import (
	"genproto/common"
	"strings"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/helper"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/pkg/errors"
)

type measurementUnitRepo struct {
	db  models.DB
	log logger.Logger
}

func NewMeasurementUnitRepo(db models.DB, log logger.Logger) repo.MeasurementI {
	return &measurementUnitRepo{
		db:  db,
		log: log,
	}
}

func (c *measurementUnitRepo) Upsert(entity *common.MeasurementUnitCopyRequest) error {

	query := `
		INSERT INTO
			"measurement_unit"
		(
			id,
			company_id,
			is_deletable,
			short_name,
			long_name,
			precision
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`

	_, err := c.db.Exec(
		query,
		entity.Id,
		entity.CompanyId,
		entity.IsDeletable,
		entity.ShortName,
		entity.LongName,
		entity.Precision,
	)
	if err != nil {
		return errors.Wrap(err, "error while upsert measurement_unit")
	}

	return nil
}

func (c *measurementUnitRepo) UpsertMany(req *common.MeasurementUnitsCopyRequest) error {

	var values = []interface{}{}

	query := `
		INSERT INTO
			"measurement_unit"
		(
			id,
			company_id,
			is_deletable,
			short_name,
			long_name,
			precision
		)
		VALUES
	`

	for _, measurement_unit := range req.MeasurementUnits {
		query += "(?, ?, ?, ?, ?, ?),"
		values = append(values,
			measurement_unit.Id,
			measurement_unit.CompanyId,
			measurement_unit.IsDeletable,
			measurement_unit.ShortName,
			measurement_unit.LongName,
			measurement_unit.Precision,
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while upsert measurement_units. Prepare")
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while upsert measurement_units. Exec")
	}

	stmt.Close()

	return nil
}

func (c *measurementUnitRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}
