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

type RegisteredUserFollowersHandler struct {
	Service * service.RegisteredUserFollowersService
}

func (handler *RegisteredUserFollowersHandler) CreateRegisteredUserFollowers(w http.ResponseWriter, r *http.Request) {
	var registeredUserFollowersDTO dto.RegisteredUserFollowersDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserFollowersDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registeredUserFollowers := model.RegisteredUserFollowers{
		ID:               uuid.UUID{},
		RegisteredUserId: registeredUserFollowersDTO.RegisteredUserId,
		FollowerUserId:   registeredUserFollowersDTO.FollowerUserId,
	}

	err = handler.Service.CreateRegisteredUserFollowers(&registeredUserFollowers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}