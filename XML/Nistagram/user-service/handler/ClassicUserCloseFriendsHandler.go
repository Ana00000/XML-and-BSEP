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
)

type ClassicUserCloseFriendsHandler struct {
	ClassicUserCloseFriendsService * service.ClassicUserCloseFriendsService
	ClassicUserFollowersService * service.ClassicUserFollowersService
}


func (handler *ClassicUserCloseFriendsHandler) CreateClassicUserCloseFriend(w http.ResponseWriter, r *http.Request) {
	var classicUserCloseFriendsDTO dto.ClassicUserCloseFriendsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserCloseFriendsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	// PROVERA DA LI SE MEDJUSOBNO PRATE

	var checkIfFollowing1 = handler.ClassicUserFollowersService.CheckIfFollowers(classicUserCloseFriendsDTO.CloseFriendUserId, classicUserCloseFriendsDTO.ClassicUserId)
	var checkIfFollowing2 = handler.ClassicUserFollowersService.CheckIfFollowers(classicUserCloseFriendsDTO.ClassicUserId, classicUserCloseFriendsDTO.CloseFriendUserId)

	if checkIfFollowing1 != true || checkIfFollowing2 != true{
		fmt.Println("Users are not following eachother")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	classicUserCloseFriends := model.ClassicUserCloseFriends{
		ID:               uuid.UUID{},
		ClassicUserId: classicUserCloseFriendsDTO.ClassicUserId,
		CloseFriendUserId:   classicUserCloseFriendsDTO.CloseFriendUserId,
	}
	err = handler.ClassicUserCloseFriendsService.CreateClassicUserCloseFriends(&classicUserCloseFriends)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}