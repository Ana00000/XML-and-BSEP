package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type AgentHandler struct {
	AgentService *service.AgentService
	UserService  *service.UserService
	Validator    *validator.Validate
	PasswordUtil *util.PasswordUtil
}

func (handler *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) {
	var agentDTO dto.AgentDTO
	if err := json.NewDecoder(r.Body).Decode(&agentDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&agentDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(agentDTO.Username) != nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(agentDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt, password := handler.PasswordUtil.GeneratePasswordWithSalt(agentDTO.Password)

	gender := model.OTHER
	switch agentDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	agentId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, agentDTO.DateOfBirth)
	agent := model.Agent{
		ClassicUser: model.ClassicUser{
			User: model.User{
				ID:          agentId,
				Username:    agentDTO.Username,
				Password:    password,
				Email:       agentDTO.Email,
				PhoneNumber: agentDTO.PhoneNumber,
				FirstName:   agentDTO.FirstName,
				LastName:    agentDTO.LastName,
				Gender:      gender,
				DateOfBirth: dateOfBirth,
				Website:     agentDTO.Website,
				Biography:   agentDTO.Biography,
				Salt: salt,
				IsConfirmed: false,
				UserType:    model.AGENT,
			},
			IsDeleted: false,
		},
		AgentRegistrationRequestId: agentDTO.AgentRegistrationRequestId,
	}

	if err := handler.AgentService.CreateAgent(&agent); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
