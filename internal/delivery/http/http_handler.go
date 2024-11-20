package http

import (
	kafka "github.com/Mubinabd/chat/internal/pkg/kafka/producer"
	s "github.com/Mubinabd/chat/internal/usecase"
	"github.com/go-redis/redis/v8"
)

type Handlers struct {
	Auth     *s.AuthUseCase
	User     *s.UserUseCase
	Group    *s.GroupUseCase
	Message  *s.MessageUseCase
	File     *s.FileUseCase
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(auth *s.AuthUseCase, user *s.UserUseCase, message *s.MessageUseCase, file *s.FileUseCase, group *s.GroupUseCase, rdb *redis.Client, pr *kafka.KafkaProducer) *Handlers {
	return &Handlers{
		Auth:     auth,
		User:     user,
		Group:    group,
		File:     file,
		Message:  message,
		RDB:      rdb,
		Producer: *pr,
	}
}
