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

type ClassicUserHandler struct {
	Service * service.ClassicUserService
}

func (handler *ClassicUserHandler) CreateClassicUser(w http.ResponseWriter, r *http.Request) {
	var classicUserDTO dto.ClassicUserDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := uuid.UUID{}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,classicUserDTO.DateOfBirth)
	classicUser := model.ClassicUser{
		RegisteredUser:       model.RegisteredUser{
			User:                        model.User{
				ID:               userId,
				Username:         classicUserDTO.Username,
				Password:         classicUserDTO.Password,
				Email:            classicUserDTO.Email,
				PhoneNumber:      classicUserDTO.PhoneNumber,
				FirstName:        classicUserDTO.FirstName,
				LastName:         classicUserDTO.LastName,
				Gender:           classicUserDTO.Gender,
				DateOfBirth:      dateOfBirth,
				Website:          classicUserDTO.Website,
				Biography:        classicUserDTO.Biography,
				//SentMessages:     nil,
				//ReceivedMessages: nil,
			},
			//Following:                   nil,
			//Followers:                   nil,
			//Campaigns:                   nil,
			//InappropriateContentRequest: nil,
		},
		//Stories:              nil,
		//StoryHighlights:      nil,
		//Posts:                nil,
		//PostCollections:      nil,
		//Activities:           nil,
		//Comments:             nil,
		IsBlocked:            false,
		UserCategory:         classicUserDTO.UserCategory,
		OfficialDocumentPath: classicUserDTO.OfficialDocumentPath,
		//SettingsId: classicUserDTO.SettingsId,
	}

	err = handler.Service.CreateClassicUser(&classicUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


