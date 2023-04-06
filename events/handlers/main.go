package handlers

import (
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/Invan2/invan_order_service/storage"
)

type EventHandler struct {
	log    logger.Logger
	strgPG storage.StorageI
}

func NewHandler(log logger.Logger, strgPG storage.StorageI) *EventHandler {
	return &EventHandler{
		log:    log,
		strgPG: strgPG,
	}
}
