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

type VerificationRequestHandler struct {
	Service * service.VerificationRequestService
}

func (handler *VerificationRequestHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request) {
	var verificationRequestDTO dto.VerificationRequestDTO
	err := json.NewDecoder(r.Body).Decode(&verificationRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	verificationRequest := model.VerificationRequest{
		ID:          			   uuid.UUID{},
		FirstName:   			   verificationRequestDTO.FirstName,
		LastName:     			   verificationRequestDTO.LastName,
		OfficialDocumentPath:      verificationRequestDTO.OfficialDocumentPath,
		RegisteredUserCategory:         verificationRequestDTO.RegisteredUserCategory,
	}

	err = handler.Service.CreateVerificationRequest(&verificationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

