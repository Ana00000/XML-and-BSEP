package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
	"time"
)

type RegisteredUserHandler struct {
	Service * service.RegisteredUserService
}

func (handler *RegisteredUserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	var registeredUserDTO dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,registeredUserDTO.DateOfBirth)
	registeredUser := model.RegisteredUser{
		User: model.User{
			ID:               uuid.UUID{},
			Username:         registeredUserDTO.Username,
			Password:         registeredUserDTO.Password,
			Email:            registeredUserDTO.Email,
			PhoneNumber:      registeredUserDTO.PhoneNumber,
			FirstName:        registeredUserDTO.FirstName,
			LastName:         registeredUserDTO.LastName,
			Gender:           registeredUserDTO.Gender,
			DateOfBirth:      dateOfBirth,
			Website:          registeredUserDTO.Website,
			Biography:        registeredUserDTO.Biography,
			//SentMessages:     nil,
			//ReceivedMessages: nil,
		},
		//Following:                   nil,
		//Followers:                   nil,
		//Campaigns:                   nil,
		//InappropriateContentRequest: nil,
	}

	err = handler.Service.CreateRegisteredUser(&registeredUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

