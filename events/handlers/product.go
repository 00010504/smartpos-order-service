package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"genproto/catalog_service"
	"genproto/common"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) UpsertProduct(ctx context.Context, event *kafka.Message) error {

	var req common.CreateProductCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create product event", logger.Any("event", req))

	tr, err := e.strgPG.WithTransaction()
	if err != nil {
		return errors.Wrap(err, "error while run transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	err = tr.Product().Upsert(&req)
	if err != nil {
		e.log.Error("error while upsert product", logger.Error(err))

		return err
	}

	return nil

}

func (e *EventHandler) CreateMultipleProducts(ctx context.Context, event *kafka.Message) error {

	var req common.CreateImportProductsModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	fmt.Println("create multiple products length", len(req.Products))

	if len(req.Products) <= 0 {
		return nil
	}

	tr, err := e.strgPG.WithTransaction()
	if err != nil {
		return errors.Wrap(err, "error while run transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()
	err = tr.Product().InsertMany(req.Products)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventHandler) BulkUpdate(ctx context.Context, event *kafka.Message) error {

	var req catalog_service.ProductBulkOperationRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create product event", logger.Any("event", req))

	tr, err := e.strgPG.WithTransaction()
	if err != nil {
		return errors.Wrap(err, "error while run transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	_, err = tr.Product().BulkUpdate(&req)
	if err != nil {
		e.log.Error("error while upsert product", logger.Error(err))
		return err
	}

	return nil
}
