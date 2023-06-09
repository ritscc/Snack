package interactor

import (
	"fmt"
	"context"

	"github.com/ritscc/Snack/domain/model"
	"github.com/ritscc/Snack/logic/bto"
)

type UserInteractor struct {
}

func (interactor *UserInteractor) CreateUser(context.Context, *bto.CreateUserRequest) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}
	return nil
}

func (interactor *UserInteractor) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	if interactor == nil {
		return nil, fmt.Errorf("UserInteractor is nil")
	}

	user := &model.User{}

	return user, nil
}

func (interactor *UserInteractor) UpdateUser(context.Context, *model.User) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}
	return nil
}

func (interactor *UserInteractor) DeleteUser(context.Context) error {
	if interactor == nil {
		return fmt.Errorf("UserInteractor is nil")
	}
	return nil
}

