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

type ClassicUserFollowersHandler struct {
	ClassicUserFollowersService * service.ClassicUserFollowersService
	UserService * service.UserService
}

func (handler *ClassicUserFollowersHandler) CreateClassicUserFollowers(w http.ResponseWriter, r *http.Request) {
	var classicUserFollowersDTO dto.ClassicUserFollowersDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserFollowersDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	classicUserFollowers := model.ClassicUserFollowers{
		ID:               uuid.UUID{},
		ClassicUserId: classicUserFollowersDTO.ClassicUserId,
		FollowerUserId:   classicUserFollowersDTO.FollowerUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollowers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ClassicUserFollowersHandler)FindAllFollowersInfoForUser(w http.ResponseWriter, r *http.Request){

	id := r.URL.Query().Get("id")
	var loginUser = handler.UserService.FindByID(uuid.MustParse(id))

	//AUTHORIZATION

	var classicUserFollowers []model.ClassicUserFollowers
	classicUserFollowers = handler.ClassicUserFollowersService.FindAllFollowersForUser(loginUser.ID)


	var users []model.User
	users = handler.UserService.FindAllFollowersInfoForUser(classicUserFollowers)

	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
	}

	w.WriteHeader(http.StatusBadRequest)
}