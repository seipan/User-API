package main

import (
	"User-API/infra/db"
	"User-API/infra/repository"
	"User-API/usecase"
	"User-API/web/handler"
	"log"
	"net/http"
)

func main() {
	db := db.NewDriver()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserhandler(userUsecase)

	handler.InitRouter(userHandler)
	srv := &http.Server{
		Addr: ":8000",
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln("Server closed with error :", err)
	}
}
