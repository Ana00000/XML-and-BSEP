package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type RegisteredUserFollowingsHandler struct {
	Service * service.RegisteredUserFollowingsService
}

func (handler *RegisteredUserFollowingsHandler) CreateRegisteredUserFollowings(w http.ResponseWriter, r *http.Request) {
	var registeredUserFollowingsDTO dto.RegisteredUserFollowingsDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserFollowingsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registeredUserFollowings := model.RegisteredUserFollowings{
		ID:               uuid.UUID{},
		RegisteredUserId: registeredUserFollowingsDTO.RegisteredUserId,
		FollowingUserId:  registeredUserFollowingsDTO.FollowingUserId,
	}

	err = handler.Service.CreateRegisteredUserFollowings(&registeredUserFollowings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}