package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Agent/dto"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"github.com/xml/XML-and-BSEP/XML/Agent/service"
	"github.com/xml/XML-and-BSEP/XML/Agent/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type AgentUserHandler struct {
	AgentUserService  *service.AgentUserService
	Validator         *validator.Validate
	AgentPasswordUtil *util.AgentPasswordUtil
}

func (handler *AgentUserHandler) CreateAgentUser(w http.ResponseWriter, r *http.Request) {
	var agentUserDTO dto.AgentUserDTO
	if err := json.NewDecoder(r.Body).Decode(&agentUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&agentUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.AgentUserService.FindAgentByUserName(agentUserDTO.Username) != nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.AgentUserService.FindAgentByEmail(agentUserDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.AgentPasswordUtil.IsValidPassword(agentUserDTO.Password)

	if validPassword {
		salt, password = handler.AgentPasswordUtil.GeneratePasswordWithSalt(agentUserDTO.Password)
	} else {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch agentUserDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	agentUserId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, agentUserDTO.DateOfBirth)

	agentUser := model.AgentUser{
		ID:          agentUserId,
		Username:    agentUserDTO.Username,
		Password:    password,
		Email:       agentUserDTO.Email,
		PhoneNumber: agentUserDTO.PhoneNumber,
		FirstName:   agentUserDTO.FirstName,
		LastName:    agentUserDTO.LastName,
		Gender:      gender,
		DateOfBirth: dateOfBirth,
		Website:     agentUserDTO.Website,
		Biography:   agentUserDTO.Biography,
		Salt:        salt,
	}

	if err := handler.AgentUserService.CreateAgentUser(&agentUser); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
