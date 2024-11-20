package kafka

import (
	"context"
	"encoding/json"

	s "github.com/Mubinabd/chat/internal/usecase"
	pb "github.com/Mubinabd/chat/internal/entity"
	"golang.org/x/exp/slog"
)

func UserRegisterHandler(u *s.AuthUseCase) func(message []byte) {
	return func(message []byte) {
		var user pb.RegisterReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.Register(context.Background(), &user)
		if err != nil {
			slog.Error("failed to register user via Kafka: %v", err)
			return
		}
		slog.Info("User registered")
	}
}
