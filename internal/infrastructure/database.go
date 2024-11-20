package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mubinabd/chat/internal/usecase"
	repo "github.com/Mubinabd/chat/internal/repository"
	"github.com/Mubinabd/chat/internal/pkg/config"
)

type Storage struct {
	Db       *sql.DB
	AuthS    usecase.AuthI
	UserS    usecase.UserI
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
		Db:       db,
		AuthS:    repo.NewAuthRepo(db),
		UserS:    repo.NewUserRepo(db),
	}, nil
}

func (s *Storage) Auth() usecase.AuthI {
	if s.AuthS == nil {
		s.AuthS = repo.NewAuthRepo(s.Db)
	}
	return s.AuthS
}

func (s *Storage) User() usecase.UserI {
	if s.UserS == nil {
		s.UserS = repo.NewUserRepo(s.Db)
	}
	return s.UserS
}
