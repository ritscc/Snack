package repository

import (
	"context"

	"github.com/ritscc/Snack/internal/domain/model"
)

type IUserRepository interface {
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
}
