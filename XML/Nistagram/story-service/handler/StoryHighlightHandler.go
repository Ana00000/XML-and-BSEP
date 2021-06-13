package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type StoryHighlightHandler struct {
	Service * service.StoryHighlightService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CRSTRYHIGHLHT0312
func (handler *StoryHighlightHandler) CreateStoryHighlight(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "CRSTRYHIGHLHT0312",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-story-highlight-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	/*
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "CRSTRYHIGHLHT0312",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}
	*/
	var storyHighlightDTO dto.StoryHighlightDTO
	err := json.NewDecoder(r.Body).Decode(&storyHighlightDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "CRSTRYHIGHLHT0312",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryHighlightDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyHighlight := model.StoryHighlight{
		ID:     uuid.UUID{},
		Title:  storyHighlightDTO.Title,
		UserId: storyHighlightDTO.UserId,
	}

	err = handler.Service.CreateStoryHighlight(&storyHighlight)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "CRSTRYHIGHLHT0312",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story highlight!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryHighlightHandler",
		"action":   "CRSTRYHIGHLHT0312",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story highlight!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
//FIDALSTRYHIGHLHTSFORUS8882
func (handler *StoryHighlightHandler) FindAllStoryHighlightsForUser(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "FIDALSTRYHIGHLHTSFORUS8882",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-story-highlights-for-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumHandler",
			"action":   "FIDALALBSTORISFORLOGGUS8293",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHighlightHandler",
			"action":   "FIDALSTRYHIGHLHTSFORUS8882",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	storyHighlights := handler.Service.FindAllStoryHighlightsForUser(uuid.MustParse(id))
	storyHighlightsJson, _ := json.Marshal(storyHighlights)
	if storyHighlightsJson != nil {
		w.WriteHeader(http.StatusOK)
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "StoryHighlightHandler",
			"action":   "FIDALSTRYHIGHLHTSFORUS8882",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all story highlights for user!")
		w.Header().Set("Content-Type", "application/json")
		w.Write(storyHighlightsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "StoryHighlightHandler",
		"action":   "FIDALSTRYHIGHLHTSFORUS8882",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding all story highlights for user!")
	w.WriteHeader(http.StatusBadRequest)
}
