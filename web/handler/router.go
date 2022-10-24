package handler

import (
	"User-API/web/middleware"
	"net/http"
)

func InitRouter(userHandler UserHandler) {
	userCreate := http.HandlerFunc(userHandler.CreateUser)
	http.Handle("/user/create", middleware.Layres(userCreate))

	userUpdate := http.HandlerFunc(userHandler.UpdateUser)
	http.Handle("/user/update", middleware.Layres(userUpdate))

	userGet := http.HandlerFunc(userHandler.GetUser)
	http.Handle("/user/get", middleware.Layres(userGet))

	userDelete := http.HandlerFunc(userHandler.DeleteUser)
	http.Handle("/user/delete", middleware.Layres(userDelete))
}
