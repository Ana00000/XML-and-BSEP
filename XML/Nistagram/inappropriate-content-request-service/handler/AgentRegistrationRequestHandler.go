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

type AgentRegistrationRequestHandler struct {
	Service * service.AgentRegistrationRequestService
}

func (handler *AgentRegistrationRequestHandler) CreateAgentRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	var agentRegistrationRequestDTO dto.AgentRegistrationRequestDTO
	err := json.NewDecoder(r.Body).Decode(&agentRegistrationRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	agentRegistrationRequest := model.AgentRegistrationRequest{
		ID:         uuid.UUID{},
		UserId:   	agentRegistrationRequestDTO.UserId,
	}

	err = handler.Service.CreateAgentRegistrationRequest(&agentRegistrationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

