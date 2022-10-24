package handler

import (
	"User-API/domain/entity"
	handler_error "User-API/error/handler"
	"User-API/usecase"
	"User-API/utils"
	"User-API/web/response"
	"encoding/json"
	"net/http"
	"strconv"
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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateErrorResponse(w, r, "method not allowed", 404, handler_error.MethodNotAllowd)
		return
	}

	newName := r.FormValue("name")
	if newName == "" {
		utils.CreateErrorResponse(w, r, "name empty", 400, nil)
		return
	}

	newId := r.FormValue("id")
	if newId == "" {
		utils.CreateErrorResponse(w, r, "id empty", 400, nil)
		return
	}

	newMail := r.FormValue("mail")
	if newMail == "" {
		utils.CreateErrorResponse(w, r, "id empty", 400, nil)
		return
	}

	newintId, err := strconv.ParseInt(newId, 10, 64)

	if err != nil {
		utils.CreateErrorResponse(w, r, "id not number", 400, nil)
		return
	}

	newUser := &entity.User{}
	newUser.Id = newintId
	newUser.Name = newName
	newUser.Mail = newMail

	user, err := uh.userUsecase.CreateUser(r.Context(), newUser)

	if err != nil {
		utils.CreateErrorResponse(w, r, "faild to createuser", 204, err)
		return
	}

	resUser := response.NewUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(resUser); err != nil {
		utils.CreateErrorResponse(w, r, "json encode error", 204, err)
		return
	}

}

func (uh userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateErrorResponse(w, r, "method not allowed", 404, handler_error.MethodNotAllowd)
		return
	}

	newName := r.FormValue("name")
	if newName == "" {
		utils.CreateErrorResponse(w, r, "name empty", 400, nil)
		return
	}

	newMail := r.FormValue("mail")
	if newMail == "" {
		utils.CreateErrorResponse(w, r, "id empty", 400, nil)
		return
	}

	newUser := &entity.User{}
	newUser.Name = newName
	newUser.Mail = newMail

	user, err := uh.userUsecase.UpdateUser(r.Context(), newUser)

	if err != nil {
		utils.CreateErrorResponse(w, r, "faild to updateuser", 204, err)
		return
	}

	resUser := response.NewUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(resUser); err != nil {
		utils.CreateErrorResponse(w, r, "json encode error", 204, err)
		return
	}
}

func (uh userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateErrorResponse(w, r, "method not allowed", 404, handler_error.MethodNotAllowd)
		return
	}

	newId := r.FormValue("id")
	if newId == "" {
		utils.CreateErrorResponse(w, r, "id empty", 400, nil)
		return
	}

	newintId, err := strconv.ParseInt(newId, 10, 64)

	if err != nil {
		utils.CreateErrorResponse(w, r, "id not number", 400, nil)
		return
	}

	user, err := uh.userUsecase.GetUser(r.Context(), newintId)

	if err != nil {
		utils.CreateErrorResponse(w, r, "faild to getuser", 204, err)
		return
	}

	resUser := response.NewUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(resUser); err != nil {
		utils.CreateErrorResponse(w, r, "json encode error", 204, err)
		return
	}
}

func (uh userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateErrorResponse(w, r, "method not allowed", 404, handler_error.MethodNotAllowd)
		return
	}

	newId := r.FormValue("id")
	if newId == "" {
		utils.CreateErrorResponse(w, r, "id empty", 400, nil)
		return
	}

	newintId, err := strconv.ParseInt(newId, 10, 64)

	if err != nil {
		utils.CreateErrorResponse(w, r, "id not number", 400, nil)
		return
	}

	err = uh.userUsecase.DeleteUser(r.Context(), newintId)

	if err != nil {
		utils.CreateErrorResponse(w, r, "faild to getuser", 204, err)
		return
	}

	utils.CreateOKResponse(w, r, "deleteuser success")

}
