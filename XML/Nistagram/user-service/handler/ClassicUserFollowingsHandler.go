package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"os"
	_ "strconv"
)

type ClassicUserFollowingsHandler struct {
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
	ClassicUserFollowersService * service.ClassicUserFollowersService
}

// CreateClassicUserFollowing KAD NEKO KLIKNE FOLLOW NEKOGA = NJEMU SE KREIRA PRVO FOLLOWING PA ONDA FOLLOWER OVOM DRUGOM
func (handler *ClassicUserFollowingsHandler) CreateClassicUserFollowing(w http.ResponseWriter, r *http.Request) {
	var classicUserFollowingDTO dto.ClassicUserFollowingsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserFollowingDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	classicUserFollowings := model.ClassicUserFollowings{
		ID:               uuid.UUID{},
		ClassicUserId: classicUserFollowingDTO.ClassicUserId,
		FollowingUserId:   classicUserFollowingDTO.FollowingUserId,
	}

	err = handler.ClassicUserFollowingsService.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	classicUserFollower := model.ClassicUserFollowers{
		ID:               uuid.UUID{},
		ClassicUserId:    classicUserFollowingDTO.FollowingUserId,
		FollowerUserId:   classicUserFollowingDTO.ClassicUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollower)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ClassicUserFollowingsHandler) AcceptFollowerRequest(w http.ResponseWriter, r *http.Request) {
	var followRequestDTO dto.FollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&followRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var followRequestForUser dto.FollowRequestForUserDTO
	reqUrlFollowRequests := fmt.Sprintf("http://%s:%s/find_request_by_classic_user_and_follower_user_ids/%s/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"),followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)
	err = getJson(reqUrlFollowRequests, &followRequestForUser)
	if err!=nil{
		fmt.Println("Wrong cast response body to FollowRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	// UPDATE REQUEST - ACCEPTED
	//var request = handler.FollowRequestService.FindFollowRequest(followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)

	reqUrlUpdate := fmt.Sprintf("http://%s:%s/accept_follow_request/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"), followRequestForUser.ID)
	jsonOrders, _ := json.Marshal(nil)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrlUpdate)
	fmt.Println(string(jsonOrders))
	resp, err := http.Post(reqUrlUpdate, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 400 {
		print("Failed creating profile settings for user")
		w.WriteHeader(http.StatusFailedDependency)
	}


	// CREATE FOLLOWER
	classicUserFollowers := model.ClassicUserFollowers{
		ID:               uuid.UUID{},
		ClassicUserId: followRequestDTO.FollowerUserId,
		FollowerUserId:   followRequestDTO.ClassicUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollowers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	// CREATE FOLLOWING
	classicUserFollowings := model.ClassicUserFollowings{
		ID:               uuid.UUID{},
		ClassicUserId:    followRequestDTO.ClassicUserId,
		FollowingUserId:   followRequestDTO.FollowerUserId,
	}

	err = handler.ClassicUserFollowingsService.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}