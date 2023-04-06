package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"genproto/common"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateCompany(ctx context.Context, event *kafka.Message) error {

	var req common.CompanyCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create company event", logger.Any("event", req))

	if err := e.strgPG.Company().Upsert(&req); err != nil {
		e.log.Error("error while upsert company", logger.Error(err), logger.Any("req", req))
	}

	e.log.Info("company shop is about to create", logger.Any("event", req.Shop))

	if req.Shop == nil {
		e.log.Error("error while upsert shop, shop == nil", logger.Any("shop", req.Shop))
	}

	if err := e.strgPG.Shop().Upsert(req.Shop); err != nil {
		e.log.Error("error while upsert shop", logger.Error(err), logger.Any("shop", req.Shop))
	}

	if req.Cashbox == nil {
		e.log.Error("error while upsert cashbox, cashbox == nil", logger.Any("cashbox", req.Cashbox))
	}

	e.log.Info("shop cashbox is about to create", logger.Any("event", req.Cashbox))

	if err := e.strgPG.Cashbox().Upsert(req.Cashbox); err != nil {
		e.log.Error("error while create cashbox", logger.Error(err), logger.Any("cashbox", req.Cashbox))
	}

	if len(req.Cheques) <= 0 {
		return errors.New("error while creating cashbox cheques")

	}

	e.log.Info("shop cheque is about to create", logger.Any("event", req.Cheques))

	if err := e.strgPG.Cheque().UpsertMany(req.Cheques); err != nil {
		e.log.Error("error while create cheque", logger.Error(err), logger.Any("cheque", req.Cheques))
	}

	e.log.Info("payment_type is about to create", logger.Any("event", req.PaymentTypes))

	if err := e.strgPG.PaymentType().UpsertMany(req.PaymentTypes); err != nil {
		e.log.Error("error while create payment_type", logger.Error(err), logger.Any("payment_type", req.PaymentTypes))
	}

	return nil

}
