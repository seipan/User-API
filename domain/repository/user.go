package repository

import (
	"context"

	"User-API/domain/entity"
)

type UserRepository interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
