package handlers

import (
	"context"
	"encoding/json"
	"genproto/catalog_service"

	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) UpsertSupplierOrder(ctx context.Context, event *kafka.Message) error {

	var request catalog_service.UpsertShopMeasurmentValueRequest

	if err := json.Unmarshal(event.Value, &request); err != nil {
		return err
	}

	e.log.Info("update mv", logger.Any("data", request))

	if err := e.strgPG.Product().UpsertShopMeasurmentValue(&request); err != nil {
		return err
	}

	return nil
}
