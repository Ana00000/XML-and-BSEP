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
)

type RegisteredUserHandler struct {
	Service * service.RegisteredUserService
}

func (handler *RegisteredUserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	var registeredUserDTO dto.RegisteredUserDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registeredUser := model.RegisteredUser{
		User : model.User{
			ID:          uuid.UUID{},
			Username:    registeredUserDTO.Username,
			Password:    registeredUserDTO.Password,
			Email:       registeredUserDTO.Email,
			PhoneNumber: registeredUserDTO.PhoneNumber,
			FirstName:   registeredUserDTO.FirstName,
			LastName:    registeredUserDTO.LastName,
			Gender:      registeredUserDTO.Gender,
			DateOfBirth: registeredUserDTO.DateOfBirth,
			Website:     registeredUserDTO.Website,
			Biography:   registeredUserDTO.Biography,
		},
	}

	err = handler.Service.CreateRegisteredUser(&registeredUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

