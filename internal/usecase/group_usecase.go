package usecase

import (
	"context"

	st "github.com/Mubinabd/chat/internal/storage/repository"
	pb "github.com/Mubinabd/chat/internal/entity"
)
type GroupUseCase struct {
	storage st.Storage

}

func NewGroupUseCase(groupRepo *st.Storage) *GroupUseCase {
    return &GroupUseCase{storage: *groupRepo}
}

func (s *GroupUseCase) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.Void, error) {
	res, err := s.storage.Group().CreateGroup(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GroupUseCase) GetAllGroups(ctx context.Context, req *pb.ListGroupsReq) (*pb.ListGroupsRes, error) {
	res, err := s.storage.Group().GetAllGroups(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GroupUseCase) DeleteGroup(ctx context.Context, req *pb.DeleteGroupReq) (*pb.Void, error) {
	res, err := s.storage.Group().DeleteGroup(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GroupUseCase) UpdateGroup(ctx context.Context, req *pb.UpdateGroupReq) (*pb.Void, error) {
	res, err := s.storage.Group().UpdateGroup(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GroupUseCase) GetGroupByID(ctx context.Context, req *pb.GetGroupByIDReq) (*pb.Group, error) {
	res, err := s.storage.Group().GetGroupByID(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GroupUseCase) AddMemberToGroup(ctx context.Context, req *pb.AddFileToGroupReq) (*pb.Void, error) {
	res, err := s.storage.Group().AddMemberToGroup(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
