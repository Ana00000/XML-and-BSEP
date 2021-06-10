package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"time"
)

type ActivityHandler struct {
	Service * service.ActivityService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var activityDTO dto.ActivityDTO
	err := json.NewDecoder(r.Body).Decode(&activityDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "CRACT467",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ActivityDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	activity := model.Activity{
		ID: uuid.UUID{},
		PostID: activityDTO.PostID,
		UserID: activityDTO.UserID,
		LikedStatus: activityDTO.LikedStatus,
		IsFavorite: activityDTO.IsFavorite,
	}

	err = handler.Service.CreateActivity(&activity)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "CRACT467",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating activity!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	activityIDJson, _ := json.Marshal(activity.ID)
	w.Write(activityIDJson)


	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ActivityHandler",
		"action":   "CRACT467",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created activity!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ActivityHandler) FindAllLikesForPost(w http.ResponseWriter, r *http.Request) {


	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllLikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "ActivityHandler",
			"action":   "FALFP468",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all likes for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "ActivityHandler",
		"action":   "FALFP468",
		"timestamp":   time.Now().String(),
	}).Error("Likes for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllDislikesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllDislikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "ActivityHandler",
			"action":   "FADFP469",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all dislikes for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "ActivityHandler",
		"action":   "FALFP469",
		"timestamp":   time.Now().String(),
	}).Error("Dislikes for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllFavoritesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllFavoritesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "ActivityHandler",
			"action":   "FAFFP470",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all favorites for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "ActivityHandler",
		"action":   "FAFFP470",
		"timestamp":   time.Now().String(),
	}).Error("Favorites for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllActivitiesForPost(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllActivitiesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "ActivityHandler",
			"action":   "FAAFP471",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all activities for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "ActivityHandler",
		"action":   "FAAFP471",
		"timestamp":   time.Now().String(),
	}).Error("Activites for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var activityDTO dto.ActivityDTO

	err := json.NewDecoder(r.Body).Decode(&activityDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ActivityDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.UpdateActivity(&activityDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Activity not updated!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ActivityHandler",
		"action":   "UPACT472",
		"timestamp":   time.Now().String(),
	}).Info("Successfully updated activity!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}