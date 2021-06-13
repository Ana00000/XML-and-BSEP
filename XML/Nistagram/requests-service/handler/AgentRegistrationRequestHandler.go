package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type AgentRegistrationRequestHandler struct {
	Service   *service.AgentRegistrationRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *AgentRegistrationRequestHandler) CreateAgentRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var agentRegistrationRequestDTO dto.AgentRegistrationRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&agentRegistrationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "AgentRegistrationRequestHandler",
			"action":    "CRAGREGREQ2010",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to AgentRegistrationRequestDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&agentRegistrationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "AgentRegistrationRequestHandler",
			"action":    "CRAGREGREQ2010",
			"timestamp": time.Now().String(),
		}).Error("AgentRegistrationRequestDTO field isn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	agentRegistrationRequest := model.AgentRegistrationRequest{
		ID:      uuid.UUID{},
		AgentId: agentRegistrationRequestDTO.AgentId,
	}

	if err := handler.Service.CreateAgentRegistrationRequest(&agentRegistrationRequest); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "AgentRegistrationRequestHandler",
			"action":    "CRAGREGREQ2010",
			"timestamp": time.Now().String(),
		}).Error("Failed creating agent registration request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "AgentRegistrationRequestHandler",
		"action":    "CRAGREGREQ2010",
		"timestamp": time.Now().String(),
	}).Info("Successfully created agent registration request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
