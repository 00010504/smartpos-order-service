package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"
	"genproto/marketing_service"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (h *EventHandler) UpsertClient(ctx context.Context, event *kafka.Message) error {

	var req marketing_service.ShortClient

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	h.log.Info("create client event", logger.Any("client", req))

	if err := h.strgPG.Client().Upsert(&req); err != nil {

		return err
	}

	return nil
}

func (h *EventHandler) DeleteClient(ctx context.Context, event *kafka.Message) error {

	var req common.RequestID

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	if err := h.strgPG.Client().Delete(&req); err != nil {

		return err
	}

	return nil
}
