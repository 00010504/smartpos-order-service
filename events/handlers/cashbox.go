package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateCashbox(ctx context.Context, event *kafka.Message) error {

	var req common.CashboxCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create cashbox event", logger.Any("event", req))

	if err := e.strgPG.Cashbox().Upsert(&req); err != nil {
		e.log.Error("error while upsert cashbox", logger.Error(err))

		return err
	}

	return nil
}
