package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type StoryTagStoriesHandler struct {
	Service * service.StoryTagStoriesService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRSTRYTGSTORIS92123
func (handler *StoryTagStoriesHandler) CreateStoryTagStories(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "CRSTRYTGSTORIS92123",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-story-tag-stories-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "CRSTRYTGSTORIS92123",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "CRSTRYTGSTORIS92123",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	var storyTagStoriesDTO dto.StoryTagStoriesDTO
	err := json.NewDecoder(r.Body).Decode(&storyTagStoriesDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "FIDALTGSFORPSTSTGPSTS9",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyTagStories := model.StoryTagStories{
		ID:         uuid.UUID{},
		TagId: storyTagStoriesDTO.TagId,
		StoryId:    storyTagStoriesDTO.StoryId,
	}

	err = handler.Service.CreateStoryTagStories(&storyTagStories)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "FIDALTGSFORPSTSTGPSTS9",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryTagStoriesHandler",
		"action":   "FIDALTGSFORPSTSTGPSTS9",
		"timestamp":   time.Now().String(),
	}).Info("Successfully added story tag for story!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFORSTRY8212
func (handler *StoryTagStoriesHandler) FindAllTagsForStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "FIDALTGSFORSTRY8212",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStory(&singleStoryDTO)

	contentsForStoriesJson, _ := json.Marshal(tags)
	w.Write(contentsForStoriesJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryTagStoriesHandler",
		"action":   "FIDALTGSFORSTRY8212",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all tags for story!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFORSORIS8223
func (handler *StoryTagStoriesHandler) FindAllTagsForStories(w http.ResponseWriter, r *http.Request) {
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryTagStoriesHandler",
			"action":   "FIDALTGSFORSORIS8223",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SingleStoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStories(singleStoriesDTO)

	contentsForStoriesJson, _ := json.Marshal(tags)
	w.Write(contentsForStoriesJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryTagStoriesHandler",
		"action":   "FIDALTGSFORSORIS8223",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all tags for stories!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDSTRYTGSTORISFORSTRYID7664
func (handler *StoryTagStoriesHandler) FindStoryTagStoriesForStoryId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	storyTagStories := handler.Service.FindStoryTagStoriesForStoryId(uuid.MustParse(id))
	storyTagStoriesJson, _ := json.Marshal(storyTagStories)
	if storyTagStoriesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "StoryTagStoriesHandler",
			"action":   "FIDALTGSFORSORIS8223",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded story tag for stories for story id!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(storyTagStoriesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "StoryTagStoriesHandler",
		"action":   "FIDALTGSFORSORIS8223",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding story tag for stories for story id!")
	w.WriteHeader(http.StatusBadRequest)
}