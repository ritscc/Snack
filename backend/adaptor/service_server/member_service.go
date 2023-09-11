package service_server

import (
	pbUserGroup "github.com/ritscc/Snack/pb/user_group/v1"
)

type MemberService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbUserGroup.UnimplementedMemberServiceServer
}

func InitMemberService() *MemberService {
	return &MemberService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}
