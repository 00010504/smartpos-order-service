package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/pkg/errors"
)

type shopRepo struct {
	db  models.DB
	log logger.Logger
}

func NewShopRepo(db models.DB, log logger.Logger) repo.ShopI {
	return &shopRepo{
		db:  db,
		log: log,
	}
}

func (s *shopRepo) Upsert(entity *common.ShopCreatedModel) error {

	query := `
		INSERT INTO
			"shop"
		(
			id,
			title,
			company_id
		)
		VALUES (
			$1,
			$2,
			$3
		) ON CONFLICT (id) DO
		UPDATE
			SET
			title = $2,
			company_id = $3;
	`

	_, err := s.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert shop")
	}

	return nil
}

func (s *shopRepo) GetById(req *common.RequestID) (*common.ShortShop, error) {

	var res common.ShortShop

	query := `
		SELECT
			"id",
			"title"
		FROM "shop"
		WHERE id = $1 AND company_id = $2 AND deleted_at = 0
	`

	err := s.db.QueryRow(query, req.Id, req.Request.CompanyId).Scan(&res.Id, &res.Name)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &res, nil
}

func (s *shopRepo) DeleteById(req *common.RequestID) error {

	query := `
		UPDATE
			"shop"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE
			id = $1,
			company_id = $2;
	`

	_, err := s.db.Exec(query, req.Id, req.Request.CompanyId)
	if err != nil {
		return errors.Wrap(err, "error while delete shop")
	}

	return nil
}
