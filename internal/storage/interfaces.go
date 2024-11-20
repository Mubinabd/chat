package repository

import (
	pb "github.com/Mubinabd/chat/internal/entity"
)

type StorageI interface {
	Auth() AuthI
	User() UserI
	File() FileI
	Group() GroupI
	Message() MessageI
}

type AuthI interface {
	Register(*pb.RegisterReq) (*pb.Void, error)
	Login(*pb.LoginReq) (*pb.User, error)
	ForgotPassword(*pb.GetByEmail) (*pb.Void, error)
	ResetPassword(*pb.ResetPassReq) (*pb.Void, error)
	SaveRefreshToken(*pb.RefToken) (*pb.Void, error)
	GetAllUsers(*pb.ListUserReq) (*pb.ListUserRes, error)
}

type UserI interface {
	GetProfile(*pb.GetById) (*pb.UserRes, error)
	EditProfile(*pb.UserRes) (*pb.UserRes, error)
	ChangePassword(*pb.ChangePasswordReq) (*pb.Void, error)
	GetSetting(*pb.GetById) (*pb.Setting, error)
	EditSetting(*pb.SettingReq) (*pb.Void, error)
	DeleteUser(*pb.GetById) (*pb.Void, error)
}

type FileI interface {
    SaveFile(*pb.SaveFile) (*pb.File, error)
}

type GroupI interface {
    CreateGroup(*pb.CreateGroupReq) (*pb.Void, error)
    GetAllGroups(*pb.ListGroupsReq) (*pb.ListGroupsRes, error)
	DeleteGroup(*pb.DeleteGroupReq) (*pb.Void, error)
	UpdateGroup(*pb.UpdateGroupReq) (*pb.Void, error)
    GetGroupByID(*pb.GetGroupByIDReq) (*pb.Group, error)
    AddMemberToGroup(*pb.AddFileToGroupReq) (*pb.Void,error)
}
type MessageI interface {
	AddMessageToGroup(*pb.MessageReq) (*pb.Void,error)
}

type WebSocketHandler struct {
    FileUC  FileI
    GroupUC GroupI
}
