package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type StoryAlbumTagStoryAlbumsHandler struct {
	Service * service.StoryAlbumTagStoryAlbumsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRSTRYALBTGSTRYALBMS7677
func (handler *StoryAlbumTagStoryAlbumsHandler) CreateStoryAlbumTagStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumTagStoryAlbumsDTO dto.StoryAlbumTagStoryAlbumsDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumTagStoryAlbumsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumTagStoryAlbumsHandler",
			"action":   "CRSTRYALBTGSTRYALBMS7677",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumTagStoryAlbumsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	storyAlbumTagStoryAlbums := model.StoryAlbumTagStoryAlbums{
		ID:           id,
		TagId:        storyAlbumTagStoryAlbumsDTO.TagId,
		StoryAlbumId: storyAlbumTagStoryAlbumsDTO.StoryAlbumId,
	}

	err = handler.Service.CreateStoryAlbumTagStoryAlbums(&storyAlbumTagStoryAlbums)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumTagStoryAlbumsHandler",
			"action":   "CRSTRYALBTGSTRYALBMS7677",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story album tag for story album!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumTagStoryAlbumsHandler",
		"action":   "CRSTRYALBTGSTRYALBMS7677",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story album tag for story album!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFORSTRYALBTGSTRYALBMS6765
func (handler *StoryAlbumTagStoryAlbumsHandler) FindAllTagsForStoryAlbumTagStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO []dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumTagStoryAlbumsHandler",
			"action":   "FIDALTGSFORSTRYALBTGSTRYALBMS6765",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list StoryAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStoryAlbumTagStoryAlbums(storyAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumTagStoryAlbumsHandler",
		"action":   "FIDALTGSFORSTRYALBTGSTRYALBMS6765",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all tags for story album tag for story album!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFORSTRYALB1980
func (handler *StoryAlbumTagStoryAlbumsHandler) FindAllTagsForStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumTagStoryAlbumsHandler",
			"action":   "FIDALTGSFORSTRYALB1980",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStoryAlbum(&storyAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumTagStoryAlbumsHandler",
		"action":   "FIDALTGSFORSTRYALB1980",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all tags for story album!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
