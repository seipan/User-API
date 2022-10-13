package repository

import (
	"User-API/domain/entity"
	"User-API/domain/repository"
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur userRepository) CreateUser(context.Context, *entity.User) (*entity.User, error) {
	return nil, nil
}

func (ur userRepository) UpdateUser(context.Context, *entity.User) (*entity.User, error) {
	return nil, nil
}

func (ur userRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (ur userRepository) DeleteUser(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}
