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

	// CHECK IF ALREADY EXISTS - IF YES THEN UPDATE TO PENDING IF NOT CREATE NEW PENDING
	var checkIfExists = handler.Service.FindFollowRequest(followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)
	if checkIfExists == nil{
		followRequest := model.FollowRequest{
			ID:          			   uuid.UUID{},
			ClassicUserId:   		   followRequestDTO.ClassicUserId,
			FollowerUserId:     	   followRequestDTO.FollowerUserId,
			FollowRequestStatus:       model.PENDING,
		}

		err = handler.Service.CreateFollowRequest(&followRequest)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}

	}else{

		err = handler.Service.UpdateFollowRequestPending(checkIfExists.ID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}

	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowRequestHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var request = handler.Service.FindById(uuid.MustParse(id))
	if request == nil{
		fmt.Println("Request not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.Service.UpdateFollowRequestRejected(uuid.MustParse(id))
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) FindAllPendingFollowerRequestsForUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var requests = handler.Service.FindAllPendingFollowerRequestsForUser(uuid.MustParse(id))
	if  requests == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	requestsJson, _ := json.Marshal(requests)
	w.Write(requestsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowRequestHandler) FindRequestById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var request = handler.Service.FindById(uuid.MustParse(id))
	if  request == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	requestJson, _ := json.Marshal(request)
	w.Write(requestJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
