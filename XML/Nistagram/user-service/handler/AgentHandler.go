package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type AgentHandler struct {
	Service * service.AgentService
}

func (handler *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) {
	var agentDTO dto.AgentDTO
	err := json.NewDecoder(r.Body).Decode(&agentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,agentDTO.DateOfBirth)
	agent := model.Agent{
		RegisteredUser:             model.RegisteredUser{
			User:                        model.User{
				ID:               uuid.UUID{},
				Username:         agentDTO.Username,
				Password:         agentDTO.Password,
				Email:            agentDTO.Email,
				PhoneNumber:      agentDTO.PhoneNumber,
				FirstName:        agentDTO.FirstName,
				LastName:         agentDTO.LastName,
				Gender:           agentDTO.Gender,
				DateOfBirth:      dateOfBirth,
				Website:          agentDTO.Website,
				Biography:        agentDTO.Biography,
				IsConfirmed: 	  false,
			},
		},
		AgentRegistrationRequestId: agentDTO.AgentRegistrationRequestId,
	}

	err = handler.Service.CreateAgent(&agent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

