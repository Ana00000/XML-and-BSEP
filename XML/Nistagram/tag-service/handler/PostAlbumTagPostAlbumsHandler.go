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

type PostAlbumTagPostAlbumsHandler struct {
	Service * service.PostAlbumTagPostAlbumsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRPOALBTGPOALBMS9832
func (handler *PostAlbumTagPostAlbumsHandler) CreatePostAlbumTagPostAlbums(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "CRPOALBTGPOALBMS9832",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-post-album-tag-post-albums-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "CRPOALBTGPOALBMS9832",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	var postAlbumTagPostAlbumsDTO dto.PostAlbumTagPostAlbumsDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumTagPostAlbumsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "CRPOALBTGPOALBMS9832",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumTagPostAlbumsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postAlbumTagPostAlbums := model.PostAlbumTagPostAlbums{
		ID:          id,
		TagId:       postAlbumTagPostAlbumsDTO.TagId,
		PostAlbumId: postAlbumTagPostAlbumsDTO.PostAlbumId,
	}

	err = handler.Service.CreatePostAlbumTagPostAlbums(&postAlbumTagPostAlbums)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "CRPOALBTGPOALBMS9832",
			"timestamp":   time.Now().String(),
		}).Error("Failed adding post album tag for post album!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumTagPostAlbumsHandler",
		"action":   "CRPOALBTGPOALBMS9832",
		"timestamp":   time.Now().String(),
	}).Info("Successfully added post album tag for post album!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFRPOALBTGPSTALBMS321
func (handler *PostAlbumTagPostAlbumsHandler) FindAllTagsForPostAlbumTagPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO []dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "FIDALTGSFRPOALBTGPSTALBMS321",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPostAlbumTagPostAlbums(postAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumTagPostAlbumsHandler",
		"action":   "FIDALTGSFRPOALBTGPSTALBMS321",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tags for album tag for post album!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFRPOALB9231
func (handler *PostAlbumTagPostAlbumsHandler) FindAllTagsForPostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumTagPostAlbumsHandler",
			"action":   "FIDALTGSFRPOALB9231",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPostAlbum(&postAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumTagPostAlbumsHandler",
		"action":   "FIDALTGSFRPOALB9231",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tags for post album!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
