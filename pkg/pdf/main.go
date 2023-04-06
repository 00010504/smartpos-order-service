package pdfmaker

import (
	"github.com/minio/minio-go/v7"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/models"
	"github.com/Invan2/invan_order_service/pkg/logger"
)

type PdFmaker struct {
	log         logger.Logger
	cfg         *config.Config
	minioClient *minio.Client
}

type PdFMakekerI interface {
	MakeOrderHTML(order *models.CreatedOrderPDFRequest) (string, error)
	HTMLtoPDF(htmlPath string, pdfName string) (string, error)
}

func NewPdfMaker(log logger.Logger, minioClient *minio.Client, config *config.Config) PdFMakekerI {

	return &PdFmaker{
		log:         log,
		minioClient: minioClient,
		cfg:         config,
	}
}
