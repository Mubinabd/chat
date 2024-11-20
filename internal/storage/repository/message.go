package repository

import (
	"database/sql"
	"errors"

	"github.com/Mubinabd/chat/internal/entity"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) AddMessageToGroup(req *entity.MessageReq) (*entity.Void, error) {
	if req.GroupID == "" || req.Message == "" {
		return &entity.Void{}, errors.New("groupID and message cannot be empty")
	}

	query := `
		INSERT INTO group_messages (group_id, message, created_at)
		VALUES ($1, $2, NOW())
	`

	_, err := r.db.Exec(query, req.GroupID, req.Message)
	if err != nil {
		return &entity.Void{}, err 
	}

	return &entity.Void{}, nil
}
