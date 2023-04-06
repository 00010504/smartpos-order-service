package services

import (
	"genproto/inventory_service"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/pkg/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPCServices struct {
	log              logger.Logger
	cfg              *config.Config
	inventoryService inventory_service.InventoryServiceClient
}

type GRPCServices interface {
	InventoryService() inventory_service.InventoryServiceClient
}

func NewGrpcServices(log logger.Logger, cfg *config.Config) (GRPCServices, error) {

	connInventory, err := grpc.Dial(cfg.InventoryService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "error while creating order service client")
	}

	res := &gRPCServices{
		log:              log,
		cfg:              cfg,
		inventoryService: inventory_service.NewInventoryServiceClient(connInventory),
	}

	return res, nil
}

func (s *gRPCServices) InventoryService() inventory_service.InventoryServiceClient {
	return s.inventoryService
}
