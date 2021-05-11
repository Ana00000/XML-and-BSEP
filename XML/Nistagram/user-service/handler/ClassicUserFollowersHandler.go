package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type ClassicUserFollowersHandler struct {
	Service * service.ClassicUserFollowersService
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

	err = handler.Service.CreateClassicUserFollowers(&classicUserFollowers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}