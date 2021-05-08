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

type AgentUserHandler struct {
	Service * service.AgentUserService
}

func (handler *AgentUserHandler) CreateAgentUser(w http.ResponseWriter, r *http.Request) {
	var agentUserDTO dto.AgentUserDTO
	err := json.NewDecoder(r.Body).Decode(&agentUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,agentUserDTO.DateOfBirth)
	agentUser := model.AgentUser{
		ID:               uuid.UUID{},
		Username:         agentUserDTO.Username,
		Password:         agentUserDTO.Password,
		Email:            agentUserDTO.Email,
		PhoneNumber:      agentUserDTO.PhoneNumber,
		FirstName:        agentUserDTO.FirstName,
		LastName:         agentUserDTO.LastName,
		Gender:           agentUserDTO.Gender,
		DateOfBirth:      dateOfBirth,
		Website:          agentUserDTO.Website,
		Biography:        agentUserDTO.Biography,
	}

	err = handler.Service.CreateAgentUser(&agentUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

