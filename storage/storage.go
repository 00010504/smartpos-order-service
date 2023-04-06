package storage

import (
	"context"
	"database/sql"

	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage/postgres"
	"github.com/Invan2/invan_order_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type repos struct {
	orderRepo           repo.OrderI
	companyRepo         repo.CompanyI
	userRepo            repo.UserI
	productRepo         repo.ProductI
	chequeRepo          repo.ChequeI
	shopRepo            repo.ShopI
	cashboxRepo         repo.CashboxI
	measurementUnitRepo repo.MeasurementI
	paymentTypeRepo     repo.PaymentTypeI
	clientRepo          repo.ClientI
}

type repoIs interface {
	Order() repo.OrderI
	Company() repo.CompanyI
	User() repo.UserI
	Product() repo.ProductI
	Cheque() repo.ChequeI
	Shop() repo.ShopI
	Cashbox() repo.CashboxI
	MeasurementUnit() repo.MeasurementI
	PaymentType() repo.PaymentTypeI
	Client() repo.ClientI
}

type storage struct {
	db  *sqlx.DB
	log logger.Logger
	repos
}

type storageTr struct {
	tr *sqlx.Tx
	repos
}

type StorageTrI interface {
	Commit() error
	Rollback() error
	repoIs
}

type StorageI interface {
	WithTransaction() (StorageTrI, error)
	repoIs
}

func NewStoragePg(log logger.Logger, db *sqlx.DB) StorageI {

	return &storage{
		db:    db,
		log:   log,
		repos: getAllRepos(log, db),
	}
}

func (s *storage) WithTransaction() (StorageTrI, error) {

	tr, err := s.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &storageTr{
		tr:    tr,
		repos: getAllRepos(s.log, tr),
	}, nil
}

func getAllRepos(log logger.Logger, db models.DB) repos {
	return repos{
		orderRepo:           postgres.NewOrderRepo(log, db),
		companyRepo:         postgres.NewCompanyRepo(log, db),
		userRepo:            postgres.NewUserRepo(log, db),
		productRepo:         postgres.NewProductRepo(log, db),
		chequeRepo:          postgres.NewChequeRepo(db, log),
		shopRepo:            postgres.NewShopRepo(db, log),
		cashboxRepo:         postgres.NewCashboxRepo(db, log),
		measurementUnitRepo: postgres.NewMeasurementUnitRepo(db, log),
		paymentTypeRepo:     postgres.NewPaymentTypeRepo(db, log),
		clientRepo:          postgres.NewClientRepo(log, db),
	}
}

func (s *storageTr) Commit() error {
	return s.tr.Commit()
}

func (s *storageTr) Rollback() error {
	return s.tr.Rollback()
}

func (r *repos) Order() repo.OrderI {
	return r.orderRepo
}

func (r *repos) Company() repo.CompanyI {
	return r.companyRepo
}

func (r *repos) User() repo.UserI {
	return r.userRepo
}

func (r *repos) Product() repo.ProductI {
	return r.productRepo
}

func (r *repos) Cheque() repo.ChequeI {
	return r.chequeRepo
}

func (r *repos) Shop() repo.ShopI {
	return r.shopRepo
}

func (r *repos) Cashbox() repo.CashboxI {
	return r.cashboxRepo
}

func (r *repos) MeasurementUnit() repo.MeasurementI {
	return r.measurementUnitRepo
}

func (r *repos) PaymentType() repo.PaymentTypeI {
	return r.paymentTypeRepo
}

func (r *repos) Client() repo.ClientI {
	return r.clientRepo
}
