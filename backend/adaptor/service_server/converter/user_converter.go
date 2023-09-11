package converter

import (
	"github.com/ritscc/Snack/domain/model"
	pbUser "github.com/ritscc/Snack/pb/user/v1"

	"github.com/ritscc/Snack/logic/dto"
)

type UserConverter struct {
}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (converter UserConverter) UserTopbUser(user *model.User) *pbUser.User {
	return &pbUser.User{
		UserId:    user.UserID,
		Username:  user.UserName,
		Nick:      user.Nick,
		RealName:  user.RealName,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Locale:    user.Locale,
		RitsEmail: user.RitsEmail,
		OwnEmail:  &user.OwnEmail,
		Comment:   user.Comment,
		Status:    pbUser.Status(user.Status),
		Services: &pbUser.Service{
			Twitter: &user.Services.Twitter,
			Github:  &user.Services.Github,
			Discord: &user.Services.Discord,
		},
	}
}

func (converter UserConverter) PbUserToUser(user *pbUser.User) *model.User {
	return &model.User{
		UserID:    user.UserId,
		UserName:  user.Username,
		Nick:      user.Nick,
		RealName:  user.RealName,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Locale:    user.Locale,
		RitsEmail: user.RitsEmail,
		OwnEmail:  *user.OwnEmail,
		Comment:   user.Comment,
		Status:    model.Status(user.Status),
		Services: &model.Service{
			Twitter: *user.Services.Twitter,
			Github:  *user.Services.Github,
			Discord: *user.Services.Discord,
		},
	}
}

func (converter UserConverter) PbCreateUserRequestToCreateUserRequest(req *pbUser.CreateUserRequest) *dto.CreateUserRequest {
	return &dto.CreateUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
