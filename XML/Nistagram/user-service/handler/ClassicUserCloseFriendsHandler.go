package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"time"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
)

type ClassicUserCloseFriendsHandler struct {
	ClassicUserCloseFriendsService * service.ClassicUserCloseFriendsService
	ClassicUserFollowersService * service.ClassicUserFollowersService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CHEKCLOFR219
func (handler *ClassicUserCloseFriendsHandler) CheckIfCloseFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	logId := vars["logId"]

	var check = handler.ClassicUserCloseFriendsService.CheckIfCloseFriend(uuid.MustParse(id), uuid.MustParse(logId))

	var returnValue = ReturnValueBool{ReturnValue: check}

	returnValueJson, _ := json.Marshal(returnValue)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserCloseFriendsHandler",
		"action":   "CHEKCLOFR219",
		"timestamp":   time.Now().String(),
	}).Info("Successfully checked close friend! Result : "+convertBoolToString(check))

	w.Write(returnValueJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertBoolToString(boolVal bool) string{
	if boolVal {
		return "true"
	} else {
		return "false"
	}
}

//CRCLOFR833
func (handler *ClassicUserCloseFriendsHandler) CreateClassicUserCloseFriend(w http.ResponseWriter, r *http.Request) {
	var classicUserCloseFriendsDTO dto.ClassicUserCloseFriendsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserCloseFriendsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserCloseFriendsHandler",
			"action":   "CRCLOFR833",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ClassicUserCloseFriendsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// PROVERA DA LI SE MEDJUSOBNO PRATE

	var checkIfFollowingFirstUser = handler.ClassicUserFollowersService.CheckIfFollowers(classicUserCloseFriendsDTO.CloseFriendUserId, classicUserCloseFriendsDTO.ClassicUserId)
	var checkIfFollowingSecondUser = handler.ClassicUserFollowersService.CheckIfFollowers(classicUserCloseFriendsDTO.ClassicUserId, classicUserCloseFriendsDTO.CloseFriendUserId)

	if checkIfFollowingFirstUser != true || checkIfFollowingSecondUser != true{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserCloseFriendsHandler",
			"action":   "CRCLOFR833",
			"timestamp":   time.Now().String(),
		}).Error("Users are not following eachother!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var allCloseFriendsForUser = handler.ClassicUserCloseFriendsService.FindAllCloseFriendsForUser(classicUserCloseFriendsDTO.ClassicUserId)
	for i:=0; i<len(allCloseFriendsForUser);i++{
		if allCloseFriendsForUser[i].CloseFriendUserId == classicUserCloseFriendsDTO.CloseFriendUserId{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ClassicUserCloseFriendsHandler",
				"action":   "CRCLOFR833",
				"timestamp":   time.Now().String(),
			}).Error("User already a close friend!")
			w.WriteHeader(http.StatusConflict)//409
			return
		}
	}

	classicUserCloseFriends := model.ClassicUserCloseFriends{
		ID:                uuid.UUID{},
		ClassicUserId:     classicUserCloseFriendsDTO.ClassicUserId,
		CloseFriendUserId: classicUserCloseFriendsDTO.CloseFriendUserId,
	}
	err = handler.ClassicUserCloseFriendsService.CreateClassicUserCloseFriends(&classicUserCloseFriends)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserCloseFriendsHandler",
			"action":   "CRCLOFR833",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating close friend for classic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserCloseFriendsHandler",
		"action":   "CRCLOFR833",
		"timestamp":   time.Now().String(),
	}).Info("Successfully creating close friend for classic user!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
