package repository

import (
	"github.com/seipan/User-API/domain/entity"
)

type UserRepository interface {
	CreateUser(string, string, string) (*entity.User, error)
	UpdateUser(string, string) (*entity.User, error)
	GetUser(string) (*entity.User, error)
	DeleteUser(string) (*entity.User, error)
}
