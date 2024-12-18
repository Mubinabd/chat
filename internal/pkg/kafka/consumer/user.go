package kafka

import (
	"context"
	"encoding/json"

	"github.com/Mubinabd/chat/internal/usecase"
	pb "github.com/Mubinabd/chat/internal/entity"
	"golang.org/x/exp/slog"
)

func UserEditProfileHandler(u *usecase.UserUseCase) func(message []byte) {
	return func(message []byte) {
		var user pb.UserRes
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.EditProfile(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit user via Kafka: %v", err)
			return
		}
		slog.Info("User updated")
	}
}

func UserEditPasswordHandler(u *usecase.UserUseCase) func(message []byte) {
	return func(message []byte) {
		var user pb.ChangePasswordReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.ChangePassword(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit password via Kafka: %v", err)
			return
		}
		slog.Info("Password updated")
	}
}

func UserEditSettingHandler(u *usecase.UserUseCase) func(message []byte) {
	return func(message []byte) {
		var user pb.SettingReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.EditSetting(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit user settings via Kafka: %v", err)
			return
		}
		slog.Info("User settings updated")
	}
}
