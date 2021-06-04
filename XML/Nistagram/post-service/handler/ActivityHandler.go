package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
)

type ActivityHandler struct {
	Service *service.ActivityService
}

func (handler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var activityDTO dto.ActivityDTO
	err := json.NewDecoder(r.Body).Decode(&activityDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activity := model.Activity{
		ID:          uuid.UUID{},
		PostID:      activityDTO.PostID,
		UserID:      activityDTO.UserID,
		LikedStatus: activityDTO.LikedStatus,
		IsFavorite:  activityDTO.IsFavorite,
	}

	err = handler.Service.CreateActivity(&activity)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	activityIDJson, _ := json.Marshal(activity.ID)
	w.Write(activityIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ActivityHandler) FindAllLikesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllLikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllDislikesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllDislikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllFavoritesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllFavoritesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllActivitiesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllActivitiesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var activityDTO dto.ActivityDTO

	err := json.NewDecoder(r.Body).Decode(&activityDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.UpdateActivity(&activityDTO)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
