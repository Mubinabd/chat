package usecase

import (
	"context"

	pb "github.com/Mubinabd/chat/internal/entity"
	st "github.com/Mubinabd/chat/internal/storage/repository"
)

type AuthUseCase struct {
	storage st.Storage
}

func NewAuthUseCase(storage *st.Storage) *AuthUseCase {
	return &AuthUseCase{
		storage: *storage,
	}
}

func (s *AuthUseCase) Register(ctx context.Context, req *pb.RegisterReq) (*pb.Void, error) {
	res, err := s.storage.Auth().Register(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthUseCase) Login(ctx context.Context, req *pb.LoginReq) (*pb.User, error) {
	res, err := s.storage.Auth().Login(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthUseCase) ForgotPassword(ctx context.Context, req *pb.GetByEmail) (*pb.Void, error) {
	res, err := s.storage.Auth().ForgotPassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthUseCase) ResetPassword(ctx context.Context, req *pb.ResetPassReq) (*pb.Void, error) {
	res, err := s.storage.Auth().ResetPassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthUseCase) SaveRefreshToken(ctx context.Context, req *pb.RefToken) (*pb.Void, error) {
	res, err := s.storage.Auth().SaveRefreshToken(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthUseCase) GetAllUsers(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserRes, error) {
	res, err := s.storage.Auth().GetAllUsers(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
