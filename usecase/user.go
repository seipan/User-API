package usecase

import (
	"context"

	"github.com/seipan/User-API/domain/entity"

	"github.com/seipan/User-API/domain/repository"
)

type UserUsecase interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
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
	user, err := ur.userRepository.CreateUser(ctx, user)
	return user, err
}

func (ur userUsecase) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := ur.userRepository.UpdateUser(ctx, user)
	return user, err
}

func (ur userUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := ur.userRepository.GetUser(ctx, id)
	return user, err
}

func (ur userUsecase) DeleteUser(ctx context.Context, id string) error {
	err := ur.userRepository.DeleteUser(ctx, id)
	return err
}
