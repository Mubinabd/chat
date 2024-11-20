package usecase

import (
	"github.com/Mubinabd/chat/internal/entity"
	st "github.com/Mubinabd/chat/internal/storage/repository"
)

type MessageUseCase struct {
	storage st.Storage
}

func NewMessageUseCase(storage *st.Storage) *MessageUseCase {
	return &MessageUseCase{
		storage: *storage,
	}
}
func (m *MessageUseCase) AddMessageToGroup(req *entity.MessageReq) (*entity.Void, error) {
	return m.storage.Message().AddMessageToGroup(req)
}
