package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/pkg/errors"
)

type cashboxRepo struct {
	db  models.DB
	log logger.Logger
}

func NewCashboxRepo(db models.DB, log logger.Logger) repo.CashboxI {
	return &cashboxRepo{
		db:  db,
		log: log,
	}
}

func (c *cashboxRepo) Upsert(entity *common.CashboxCreatedModel) error {

	query := `
		INSERT INTO
			"cashbox"
		(
			id,
			shop_id,
			company_id,
			title,
			cheque_id,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) ON CONFLICT (id) DO
			UPDATE
				SET
				title = $4,
				cheque_id = $5;
			
	`

	_, err := c.db.Exec(
		query,
		entity.Id,
		entity.ShopId,
		entity.Request.CompanyId,
		entity.Title,
		entity.ChequeId,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert cashbox")
	}

	return nil
}

func (c *cashboxRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}
