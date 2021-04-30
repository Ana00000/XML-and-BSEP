package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type ActivityHandler struct {
	Service * service.ActivityService
}

func (handler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var activityDTO dto.ActivityDTO
	err := json.NewDecoder(r.Body).Decode(&activityDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activity := model.Activity{
		ID: uuid.UUID{},
		PostID: activityDTO.PostID,
		// Post
		UserID: activityDTO.PostID,
		Liked: activityDTO.Liked,
		IsFavorite: activityDTO.IsFavorite,
	}

	err = handler.Service.CreateActivity(&activity)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}