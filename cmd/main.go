package main

import (
	"context"
	"fmt"
	"genproto/order_service"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Invan2/invan_order_service/config"
	"github.com/Invan2/invan_order_service/events"

	"github.com/Invan2/invan_order_service/pkg/logger"
	pdfmaker "github.com/Invan2/invan_order_service/pkg/pdf"
	"github.com/Invan2/invan_order_service/services/listeners"
	"github.com/Invan2/invan_order_service/storage"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.Load()
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	log.Info("config", logger.Any("config", cfg), logger.Any("env", os.Environ()))

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	postgresURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	fmt.Println("postgresURL", postgresURL)
	psqlConn, err := sqlx.Connect("postgres", postgresURL)
	if err != nil {
		log.Error("poostgres", logger.Error(err))
		return
	}

	defer psqlConn.Close()

	storage := storage.NewStoragePg(log, psqlConn)

	conf := kafka.ConfigMap{
		"bootstrap.servers":                     cfg.KafkaUrl,
		"group.id":                              config.ConsumerGroupID,
		"auto.offset.reset":                     "earliest",
		"go.events.channel.size":                1000000,
		"socket.keepalive.enable":               true,
		"metadata.max.age.ms":                   900000,
		"metadata.request.timeout.ms":           30000,
		"retries":                               1000000,
		"message.timeout.ms":                    300000,
		"socket.timeout.ms":                     30000,
		"max.in.flight.requests.per.connection": 5,
		"heartbeat.interval.ms":                 3000,
		"enable.idempotence":                    true,
	}

	log.Info("kafka config", logger.Any("config", conf))

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Error("error while creating producer")
		return
	}

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Error("error while creating consumer", logger.Error(err))
		return
	}

	pubsubServer, err := events.NewPubSubServer(log, producer, consumer, storage)
	if err != nil {
		log.Fatal("error creating pubSubServer", logger.Error(err))
		return
	}

	server := grpc.NewServer()

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Secure: true,
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretKey, ""),
	})
	if err != nil {
		log.Info("minio", logger.Error(err))
		return
	}

	pdfMaker := pdfmaker.NewPdfMaker(log, minioClient, &cfg)

	listeners, err := listeners.NewOrderService(log, pubsubServer, storage, pdfMaker, &cfg)
	if err != nil {
		log.Info("new listeners", logger.Error(err))
		return
	}

	order_service.RegisterOrderServiceServer(server, listeners)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", cfg.HttpHost, cfg.HttpPort))
	if err != nil {
		log.Error("http", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
		if err := pubsubServer.Shutdown(); err != nil {
			log.Error("error while shutdown pub sub server")
			return
		}
	}()

	go func() {
		if err := pubsubServer.Run(ctx); err != nil {
			log.Error("pubsubServer.Run, error: ", logger.Error(err))
			return
		}
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatal("serve", logger.Error(err))
		return
	}
}
