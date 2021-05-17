package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	requestsDTO "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	requestsService "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
)

type ClassicUserFollowingsHandler struct {
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
	ClassicUserFollowersService * service.ClassicUserFollowersService
	FollowRequestService * requestsService.FollowRequestService
}

//KAD NEKO KLIKNE FOLLOW NEKOGA = NJEMU SE KREIRA PRVO FOLLOWING PA ONDA FOLLOWER OVOM DRUGOM
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
	var followRequestDTO requestsDTO.FollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&followRequestDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// UPDATE REQUEST - ACCEPTED
	id := r.URL.Query().Get("id")

	var request = handler.FollowRequestService.FindById(uuid.MustParse(id))
	if request == nil{
		fmt.Println("Request not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.FollowRequestService.UpdateFollowRequestAccepted(uuid.MustParse(id))

	// CREATE FOLLOWER
	classicUserFollowers := model.ClassicUserFollowers{
		ID:               uuid.UUID{},
		ClassicUserId: followRequestDTO.ClassicUserId,
		FollowerUserId:   followRequestDTO.FollowerUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollowers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	// CREATE FOLLOWING
	classicUserFollowings := model.ClassicUserFollowings{
		ID:               uuid.UUID{},
		ClassicUserId:    followRequestDTO.FollowerUserId,
		FollowingUserId:   followRequestDTO.ClassicUserId,
	}

	err = handler.ClassicUserFollowingsService.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
