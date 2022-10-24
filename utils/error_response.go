package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status int
	Result string
}

type StatusOKResponse struct {
	Status int
	Result string
}

func CreateErrorResponse(w http.ResponseWriter, r *http.Request, errormessage string, status int, err error) {
	log.Println(err)
	errResponse := ErrorResponse{status, errormessage}

	res, err := json.Marshal(errResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateOKResponse(w http.ResponseWriter, r *http.Request, statusmessage string) {
	okResponse := StatusOKResponse{http.StatusOK, statusmessage}

	res, err := json.Marshal(okResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
