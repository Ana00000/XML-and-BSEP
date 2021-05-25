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

type InappropriateContentRequestHandler struct {
	Service * service.InappropriateContentRequestService
}

func (handler *InappropriateContentRequestHandler) CreateInappropriateContentRequest(w http.ResponseWriter, r *http.Request) {
	var inappropriateContentRequestDTO dto.InappropriateContentRequestDTO
	err := json.NewDecoder(r.Body).Decode(&inappropriateContentRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inappropriateContentRequest := model.InappropriateContentRequest{
		ID:          uuid.UUID{},
		Note: 		 inappropriateContentRequestDTO.Note,
		UserId:      inappropriateContentRequestDTO.UserId,
	}

	err = handler.Service.CreateInappropriateContentRequest(&inappropriateContentRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

