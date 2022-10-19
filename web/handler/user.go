package handler

import (
	"User-API/usecase"
	"net/http"
)

type UserHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserhandler(uu usecase.UserUsecase) UserHandler {
	return userHandler{
		userUsecase: uu,
	}
}

func (uh userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
}

func (uh userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (uh userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (uh userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
