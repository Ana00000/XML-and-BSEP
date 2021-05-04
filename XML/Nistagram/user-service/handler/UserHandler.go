package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
	"time"
)

type UserHandler struct {
	Service * service.UserService
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,userDTO.DateOfBirth)
	user := model.User{
		ID:               uuid.UUID{},
		Username:         userDTO.Username,
		Password:         userDTO.Password,
		Email:            userDTO.Email,
		PhoneNumber:      userDTO.PhoneNumber,
		FirstName:        userDTO.FirstName,
		LastName:         userDTO.LastName,
		Gender:           userDTO.Gender,
		DateOfBirth:      dateOfBirth,
		Website:          userDTO.Website,
		Biography:        userDTO.Biography,
		//SentMessages:     nil,
		//ReceivedMessages: nil,
	}

	err = handler.Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

