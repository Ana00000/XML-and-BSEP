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

type ClassicUserFollowingsHandler struct {
	Service * service.ClassicUserFollowingsService
}

func (handler *ClassicUserFollowingsHandler) CreateClassicUserFollowings(w http.ResponseWriter, r *http.Request) {
	var classicUserFollowingsDTO dto.ClassicUserFollowingsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserFollowingsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	classicUserFollowings := model.ClassicUserFollowings{
		ID:               uuid.UUID{},
		ClassicUserId: classicUserFollowingsDTO.ClassicUserId,
		FollowingUserId:  classicUserFollowingsDTO.FollowingUserId,
	}

	err = handler.Service.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}