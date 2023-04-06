package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateMeasurementUnit(ctx context.Context, event *kafka.Message) error {

	var req common.MeasurementUnitCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create measurement_unit event", logger.Any("event", req))

	if err := e.strgPG.MeasurementUnit().Upsert(&req); err != nil {
		e.log.Info(err.Error(), logger.Any("event", req))

		return err
	}

	return nil
}

func (e *EventHandler) CreateMeasurementUnits(ctx context.Context, event *kafka.Message) error {

	var req common.MeasurementUnitsCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create measurement_units event", logger.Any("event", req))

	if err := e.strgPG.MeasurementUnit().UpsertMany(&req); err != nil {
		e.log.Info(err.Error(), logger.Any("event", req))

		return err
	}

	return nil
}
