package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mubinabd/chat/internal/pkg/config"
	repository "github.com/Mubinabd/chat/internal/storage"
)

type Storage struct {
	DB       *sql.DB
	AuthS    repository.AuthI
	UserS    repository.UserI
	GroupS   repository.GroupI
	MessageS repository.MessageI
	FileS    repository.FileI
}

func New(cfg *config.Config) (*Storage, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("can't connect to db: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB:       db,
		AuthS:    NewAuthRepo(db),
		UserS:    NewUserRepo(db),
		GroupS:   NewGroupRepo(db),
		MessageS: NewMessageRepo(db),
		FileS:    NewFileRepo(db),
	}, nil
}

func (s *Storage) Group() repository.GroupI {
	if s.GroupS == nil {
		s.GroupS = NewGroupRepo(s.DB)
	}
	return s.GroupS
}

func (s *Storage) Message() repository.MessageI {
	if s.MessageS == nil {
		s.MessageS = NewMessageRepo(s.DB)
	}
	return s.MessageS
}
func (s *Storage) File() repository.FileI {
	if s.FileS == nil {
		s.FileS = NewFileRepo(s.DB)
	}
	return s.FileS
}
func (s *Storage) Auth() repository.AuthI {
	if s.AuthS == nil {
		s.AuthS = NewAuthRepo(s.DB)
	}
	return s.AuthS
}

func (s *Storage) User() repository.UserI {
	if s.UserS == nil {
		s.UserS = NewUserRepo(s.DB)
	}
	return s.UserS
}
