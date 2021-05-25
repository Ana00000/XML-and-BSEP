package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
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
		ID:      uuid.UUID{},
		AgentId: agentRegistrationRequestDTO.AgentId,
	}

	err = handler.Service.CreateAgentRegistrationRequest(&agentRegistrationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

