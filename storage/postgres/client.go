package postgres

import (
	"database/sql"
	"genproto/common"
	"genproto/marketing_service"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
)

type clientRepo struct {
	log logger.Logger
	db  models.DB
}

func NewClientRepo(log logger.Logger, db models.DB) repo.ClientI {
	return &clientRepo{
		log: log,
		db:  db,
	}
}

func (c *clientRepo) Upsert(in *marketing_service.ShortClient) error {

	query := `
		INSERT INTO "client" (
			"id",
			"company_id",
			"first_name",
			"last_name",
			"phone_number"
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		) ON CONFLICT (id) DO
		UPDATE
			SET
			"first_name" = $3,
			"last_name" = $4,
			"phone_number" = $5
	`

	res, err := c.db.Exec(query,
		in.Id,
		in.CompanyId,
		in.FirstName,
		in.LastName,
		in.PhoneNumber)
	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (c *clientRepo) Delete(in *common.RequestID) error {

	query := `
		UPDATE "client" SET 
			deleted_at=extract(epoch from now())::bigint
		) WHERE id=$1 AND company_id=$2

	`

	res, err := c.db.Exec(query, in.Id, in.Request.CompanyId)
	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
