package usecase

import (
	"context"

	st "github.com/Mubinabd/chat/internal/storage/repository"
	pb "github.com/Mubinabd/chat/internal/entity"
)

type UserUseCase struct {
	storage st.Storage
}

func NewUserUseCase(storage *st.Storage) *UserUseCase {
	return &UserUseCase{
		storage: *storage,
	}
}

func (s *UserUseCase) GetProfile(ctx context.Context, req *pb.GetById) (*pb.UserRes, error) {
	res, err := s.storage.User().GetProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserUseCase) EditProfile(ctx context.Context, req *pb.UserRes) (*pb.UserRes, error) {
	res, err := s.storage.User().EditProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserUseCase) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error) {
	res, err := s.storage.User().ChangePassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserUseCase) GetSetting(ctx context.Context, req *pb.GetById) (*pb.Setting, error) {
	res, err := s.storage.User().GetSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserUseCase) EditSetting(ctx context.Context, req *pb.SettingReq) (*pb.Void, error) {
	res, err := s.storage.User().EditSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserUseCase) DeleteUser(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
