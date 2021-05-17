package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"net/http"
	_ "strconv"
)

type FollowRequestHandler struct {
	Service * service.FollowRequestService
}

func (handler *FollowRequestHandler) CreateFollowRequest(w http.ResponseWriter, r *http.Request) {
	var followRequestDTO dto.FollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&followRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followRequest := model.FollowRequest{
		ID:          			   uuid.UUID{},
		ClassicUserId:   			   followRequestDTO.ClassicUserId,
		FollowerUserId:     			   followRequestDTO.FollowerUserId,
		FollowRequestStatus:         model.PENDING,
	}

	err = handler.Service.CreateFollowRequest(&followRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

