package http

import (
	s "github.com/Mubinabd/chat/internal/service"
	kafka "github.com/Mubinabd/chat/internal/pkg/kafka/producer"
	"github.com/go-redis/redis/v8"
)

type Handlers struct {
	Auth     *s.AuthService
	User     *s.UserService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(auth *s.AuthService, user *s.UserService, rdb *redis.Client, pr *kafka.KafkaProducer) *Handlers {
	return &Handlers{
		Auth:     auth,
		User:     user,
		RDB:      rdb,
		Producer: *pr,
	}
}
