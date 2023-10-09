package interactor

import (
	"context"
	"crypto/sha512"
	"fmt"

	"github.com/ritscc/Snack/internal/domain/model"
	"github.com/ritscc/Snack/internal/usecase/dto"
)

type UserInteractor struct {
	Users   []*model.User
	Counter int64
}

func NewUserInteractor() *UserInteractor {
	return &UserInteractor{}
}

func (interactor *UserInteractor) CreateUser(ctx context.Context, req *dto.CreateUserRequest) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}

	newUser := &model.User{
		UserID:    interactor.Counter,
		UserName:  req.Username,
		RitsEmail: req.Email,
		Password:  sha512.Sum512([]byte(req.Password)),
	}

	interactor.Users = append(interactor.Users, newUser)

	interactor.Counter++

	return nil
}

func (interactor *UserInteractor) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	if interactor == nil {
		return nil, fmt.Errorf("UserInteractor is nil")
	}

	user := &model.User{}

	for _, u := range interactor.Users {
		if userID == u.UserID {
			user = u
			break
		}
	}

	return user, nil
}

func (interactor *UserInteractor) UpdateUser(ctx context.Context, user *model.User) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}

	for _, u := range interactor.Users {
		if user.UserID == u.UserID {
			u = user
			return nil
		}
	}

	return fmt.Errorf("not such a user found")
}

func (interactor *UserInteractor) DeleteUser(context.Context) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}

	// Get from jwt user token

	return nil
}
