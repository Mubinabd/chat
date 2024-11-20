package app

import (
	"context"

	a "github.com/Mubinabd/chat/internal/delivery"
	handler "github.com/Mubinabd/chat/internal/delivery/http"
	"github.com/Mubinabd/chat/internal/pkg/config"
	kafka "github.com/Mubinabd/chat/internal/pkg/kafka/consumer"
	prd "github.com/Mubinabd/chat/internal/pkg/kafka/producer"
	s "github.com/Mubinabd/chat/internal/usecase"
	"github.com/go-redis/redis/v8"
	db "github.com/Mubinabd/chat/internal/storage/repository"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {
	db, err := db.New(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}
	defer db.DB.Close()
	slog.Info("Connected to Postgres")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis_auth:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
	}
	slog.Info("Connected to Redis")

	authUseCase := s.NewAuthUseCase(db)
	userUseCase := s.NewUserUseCase(db)
	messageUseCase := s.NewMessageUseCase(db)
	groupUseCase := s.NewGroupUseCase(db)
	fileUseCase := s.NewFileUseCase(db)

	// Kafka
	brokers := []string{"kafka_auth:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		slog.Error("Failed to create Kafka producer:", err)
		return
	}

	Reader(brokers, cm, authUseCase, userUseCase)

	// HTTP Server
	h := handler.NewHandler(authUseCase, userUseCase, messageUseCase, fileUseCase, groupUseCase, rdb, &pr)

	router := a.NewGin(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.GRPCPort); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.GRPCPort)
}
