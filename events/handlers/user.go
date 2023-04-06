package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) UpsertUser(ctx context.Context, event *kafka.Message) error {

	var req common.UserCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("user create", logger.Any("event", req))

	if err := e.strgPG.User().Upsert(&req); err != nil {
		e.log.Info(err.Error(), logger.Any("event", req))

		return err
	}

	return nil
}
