package service_server

import (
	"context"

	"github.com/ritscc/Snack/internal/adaptor/service_server/converter"
	"github.com/ritscc/Snack/internal/usecase/interactor"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	pbUser "github.com/ritscc/Snack/pb/user/v1"
)

type UserService struct {
	interactor *interactor.UserInteractor
	pbUser.UnimplementedUserServiceServer
}

func InitUserService() *UserService {
	return &UserService{
		interactor: interactor.NewUserInteractor(),
	}
}

func (service *UserService) CreateUser(ctx context.Context, pbReq *pbUser.CreateUserRequest) (*emptypb.Empty, error) {
	if service == nil {
		return nil, status.Errorf(codes.Internal, "UserService is nil")
	}

	req := converter.PbCreateUserRequestToCreateUserRequest(pbReq)
	service.interactor.CreateUser(ctx, req)
	res := new(emptypb.Empty)

	return res, nil
}

func (service *UserService) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.User, error) {
	if service == nil {
		return nil, status.Errorf(codes.Internal, "UserService is nil")
	}

	user, err := service.interactor.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "err")
	}

	pbUser := converter.UserTopbUser(user)

	return pbUser, nil
}

func (service *UserService) UpdateUser(ctx context.Context, pbUser *pbUser.User) (*emptypb.Empty, error) {
	if service == nil {
		return nil, status.Errorf(codes.Internal, "UserService is nil")
	}

	user := converter.PbUserToUser(pbUser)
	if err := service.interactor.UpdateUser(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "err")
	}
	res := new(emptypb.Empty)

	return res, nil
}

func (service *UserService) DeleteUser(ctx context.Context, e *emptypb.Empty) (*emptypb.Empty, error) {
	if service == nil {
		return nil, status.Errorf(codes.Internal, "UserService is nil")
	}

	service.interactor.DeleteUser(ctx)

	res := new(emptypb.Empty)

	return res, nil
}
