package service_server

import (
	"context"

	pbUserGroup "github.com/ritscc/Snack/pb/user_group/v1"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserGroupService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbUserGroup.UnimplementedUserGroupServiceServer
}

func InitUserGroupService() *UserGroupService {
	return &UserGroupService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}

func (*UserGroupService) CreateUserGroup(context.Context, *pbUserGroup.CreateUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserGroup not implemented")
}
func (*UserGroupService) GetUserGroup(context.Context, *pbUserGroup.UserGroupID) (*pbUserGroup.UserGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroup not implemented")
}
func (*UserGroupService) UpdateUserGroup(context.Context, *pbUserGroup.UpdateUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (*UserGroupService) DeleteUserGroup(context.Context, *pbUserGroup.UserGroupID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroup not implemented")
}
