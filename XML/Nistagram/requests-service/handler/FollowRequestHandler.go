package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type FollowRequestHandler struct {
	Service   *service.FollowRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *FollowRequestHandler) CreateFollowRequest(w http.ResponseWriter, r *http.Request) {
	var followRequestDTO dto.FollowRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&followRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "CREFOLLOWREQ6319",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to FollowRequestDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&followRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "CREFOLLOWREQ6319",
			"timestamp": time.Now().String(),
		}).Error("FollowRequestDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	// CHECK IF ALREADY EXISTS - IF YES THEN UPDATE TO PENDING IF NOT CREATE NEW PENDING
	var checkIfExists = handler.Service.FindFollowRequest(followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)
	if checkIfExists == nil {
		followRequest := model.FollowRequest{
			ID:                  uuid.UUID{},
			ClassicUserId:       followRequestDTO.ClassicUserId,
			FollowerUserId:      followRequestDTO.FollowerUserId,
			FollowRequestStatus: model.PENDING,
		}

		if err := handler.Service.CreateFollowRequest(&followRequest); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "FollowRequestHandler",
				"action":    "CREFOLLOWREQ6319",
				"timestamp": time.Now().String(),
			}).Error("Failed creating follow request!")
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

	} else {

		if err := handler.Service.UpdateFollowRequestPending(checkIfExists.ID); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "FollowRequestHandler",
				"action":    "CREFOLLOWREQ6319",
				"timestamp": time.Now().String(),
			}).Error("Failed updating follow request!")
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "CREFOLLOWREQ6319",
		"timestamp": time.Now().String(),
	}).Info("Successfully create follow request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowRequestHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var request = handler.Service.FindById(uuid.MustParse(id))
	if request == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "REJFOLLOWREQ4939",
			"timestamp": time.Now().String(),
		}).Error("Reject follow request not found!")
		fmt.Println("Reject follow request not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.Service.UpdateFollowRequestRejected(uuid.MustParse(id))
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "REJFOLLOWREQ4939",
		"timestamp": time.Now().String(),
	}).Info("Successfully created reject follow request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) FindFollowRequestByIDsClassicUserAndHisFollower(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classicUserId := vars["classicUserID"]
	followerUserId := vars["followerUserID"]
	var request = handler.Service.FindFollowRequest(uuid.MustParse(classicUserId), uuid.MustParse(followerUserId))
	if request == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "FIDFOLLREQBYIDCLASUSANDHISFOLL3333",
			"timestamp": time.Now().String(),
		}).Error("Follow request by IDs classic user and his follower not found!")
		fmt.Println("Follow request by IDs classic user and his follower not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	var requestForJson = *request
	fmt.Println(requestForJson.ClassicUserId.String() + " " + requestForJson.FollowerUserId.String())
	requestsJson, _ := json.Marshal(convertFollowRequestToFollowRequestForUserDTOs(requestForJson))
	w.Write(requestsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "FIDFOLLREQBYIDCLASUSANDHISFOLL3333",
		"timestamp": time.Now().String(),
	}).Info("Successfully found follow request by IDs classic user and his follower!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) UpdateFollowRequestToAccepted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestID"]
	if err := handler.Service.UpdateFollowRequestAccepted(uuid.MustParse(requestId)); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "UPDFOLLOWREQTOACCEP7710",
			"timestamp": time.Now().String(),
		}).Error("Fail to update follow request to accept!")
		fmt.Println("Fail to update follow request to accept")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "UPDFOLLOWREQTOACCEP7710",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated follow request to accepted!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) FindAllPendingFollowerRequestsForUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var requests = handler.Service.FindAllPendingFollowerRequestsForUser(uuid.MustParse(id))

	requestsJson, _ := json.Marshal(requests)
	w.Write(requestsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "FIDALLPENFOLLOWERREQFORUS6700",
		"timestamp": time.Now().String(),
	}).Info("Successfully found all pending follower requests for user!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowRequestHandler) FindRequestById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var request = handler.Service.FindById(uuid.MustParse(id))
	if request == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "FIDREQBYID2431",
			"timestamp": time.Now().String(),
		}).Error("Request by id not found!")
		fmt.Println("Request by id not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	requestJson, _ := json.Marshal(request)
	w.Write(requestJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "FIDREQBYID2431",
		"timestamp": time.Now().String(),
	}).Info("Successfully found request by id!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowRequestHandler) FindAllFollowerRequestsForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]

	var requests = handler.Service.FindAllFollowerRequestsForUser(uuid.MustParse(userId))
	if requests == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "FollowRequestHandler",
			"action":    "FIDALLFOLLREQFORUS2491",
			"timestamp": time.Now().String(),
		}).Error("All follower requests for user not found!")
		fmt.Println("All follower requests for user not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	requestsJson, _ := json.Marshal(convertListFollowRequestsToListFollowRequestForUserDTOs(requests))
	w.Write(requestsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "FollowRequestHandler",
		"action":    "FIDALLFOLLREQFORUS2491",
		"timestamp": time.Now().String(),
	}).Info("Successfully found all follower requests for user!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func convertListFollowRequestsToListFollowRequestForUserDTOs(followerRequests []model.FollowRequest) []dto.FollowRequestForUserDTO {
	var followerRequestForUserDTOs []dto.FollowRequestForUserDTO
	for i := 0; i < len(followerRequests); i++ {
		var followRequestForUserDTO = convertFollowRequestToFollowRequestForUserDTOs(followerRequests[i])
		followerRequestForUserDTOs = append(followerRequestForUserDTOs, followRequestForUserDTO)
	}
	return followerRequestForUserDTOs
}

func convertFollowRequestToFollowRequestForUserDTOs(followerRequest model.FollowRequest) dto.FollowRequestForUserDTO {
	followRequestStatus := ""
	if followerRequest.FollowRequestStatus == model.PENDING {
		followRequestStatus = "PENDING"
	} else if followerRequest.FollowRequestStatus == model.ACCEPTED {
		followRequestStatus = "ACCEPTED"
	} else if followerRequest.FollowRequestStatus == model.REJECT {
		followRequestStatus = "REJECT"
	}
	var followRequestForUserDTO = dto.FollowRequestForUserDTO{
		ID:                  followerRequest.ID,
		ClassicUserId:       followerRequest.ClassicUserId,
		FollowerUserId:      followerRequest.FollowerUserId,
		FollowRequestStatus: followRequestStatus,
	}
	return followRequestForUserDTO
}
