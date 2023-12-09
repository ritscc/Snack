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
	//TODO: test
	return &model.User{
		UserID:    id,
		UserName:  "Ritsumei Tarou",
		Nick:      "USA",
		RealName:  "Ritsumei Tarou",
		Password:  [64]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		Avatar:    "user.png",
		Role:      "admin",
		Locale:    "ja",
		RitsEmail: "is0000aa@ed.ritsumei.ac.jp",
		OwnEmail:  "ritsumei@gmail.com",
		Comment:   "Hello, World!",
		Status:    0,
		Services: &model.Service{
			Twitter: "https://twitter.com/ritscc",
			Github:  "https://github/ritscc",
			Discord: "https://discord.gg/ritscc",
		},
	}, nil
}
