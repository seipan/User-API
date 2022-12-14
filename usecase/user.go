package usecase

import (
	"context"

	"User-API/domain/entity"
	"User-API/error/usecase"

	"User-API/domain/repository"
)

type UserUsecase interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return userUsecase{
		userRepository: ur,
	}
}

func (ur userUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase.NameEmptyError
	}
	if user.Mail == "" {
		return nil, usecase.MailEmptyError
	}
	user, err := ur.userRepository.CreateUser(ctx, user)
	return user, err
}

func (ur userUsecase) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase.NameEmptyError
	}
	if user.Mail == "" {
		return nil, usecase.MailEmptyError
	}

	user, err := ur.userRepository.UpdateUser(ctx, user)
	return user, err
}

func (ur userUsecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := ur.userRepository.GetUser(ctx, id)
	return user, err
}

func (ur userUsecase) DeleteUser(ctx context.Context, id int64) error {
	err := ur.userRepository.DeleteUser(ctx, id)
	return err
}
