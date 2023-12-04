package repository

import (
	"context"

	"github.com/ritscc/Snack/internal/domain/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
