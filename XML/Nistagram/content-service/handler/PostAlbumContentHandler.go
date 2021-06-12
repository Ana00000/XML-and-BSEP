package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type PostAlbumContentHandler struct {
	Service * service.PostAlbumContentService
	ContentService * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

var pathPostAlbumGlobal = ""

func (handler *PostAlbumContentHandler) CreatePostAlbumContent(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-post-album-content-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	var postAlbumContentDTO dto.PostAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch postAlbumContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	postAlbumContent := model.PostAlbumContent{
		Content: model.Content{
			ID:   id,
			Path: pathPostAlbumGlobal,
			Type: contentType,
		},
		PostAlbumId: postAlbumContentDTO.PostAlbumId,
	}

	err = handler.Service.CreatePostAlbumContent(&postAlbumContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post album content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.ContentService.CreateContent(&postAlbumContent.Content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "CRPOALCOL998",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumContentHandler",
		"action":   "CRPOALCOL998",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostAlbumFile")
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "UPL887",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find the file!")
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"),  "*" + hand.Filename)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "UPL887",
			"timestamp":   time.Now().String(),
		}).Error("Failed to create temporary file!")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "UPL887",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		return
	}
	tempFile.Write(fileBytes)

	pathPostAlbumGlobal = tempFile.Name()[20:len(tempFile.Name())]

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumContentHandler",
		"action":   "UPL887",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func (handler *PostAlbumContentHandler) FindAllContentsForPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO []dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "FIALCOFOPOALO555",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForPostAlbums = convertListPostAlbumContentToListPostAlbumContentDTO(handler.Service.FindAllContentsForPostAlbums(postAlbumFullDTO))

	contentsForPostAlbumsJson, _ := json.Marshal(contentsForPostAlbums)
	w.Write(contentsForPostAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumContentHandler",
		"action":   "FIALCOFOPOALO555",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for post albums!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumContentHandler) FindAllContentsForPostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAlbumContentHandler",
			"action":   "FIALCOFOPOALO673",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForPostAlbums = convertListPostAlbumContentToListPostAlbumContentDTO(handler.Service.FindAllContentsForPostAlbum(&postAlbumFullDTO))

	contentsForPostAlbumsJson, _ := json.Marshal(contentsForPostAlbums)
	w.Write(contentsForPostAlbumsJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostAlbumContentHandler",
		"action":   "FIALCOFOPOALO673",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for post album!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertPostAlbumContentToPostAlbumContentDTO(postAlbumContent model.PostAlbumContent) dto.PostAlbumContentFullDTO{
	postAlbumContentType :=""
	if postAlbumContent.Type==model.PICTURE{
		postAlbumContentType="PICTURE"
	}else if postAlbumContent.Type==model.VIDEO{
		postAlbumContentType="VIDEO"
	}
	var postAlbumContentDTO = dto.PostAlbumContentFullDTO{
		ID:           postAlbumContent.ID,
		Path:         postAlbumContent.Path,
		Type:         postAlbumContentType,
		PostAlbumId: postAlbumContent.PostAlbumId,
	}
	return postAlbumContentDTO
}

func convertListPostAlbumContentToListPostAlbumContentDTO(postAlbumContents []model.PostAlbumContent) []dto.PostAlbumContentFullDTO {
	var postAlbumContentsDTO []dto.PostAlbumContentFullDTO
	for i := 0; i < len(postAlbumContents); i++ {
		postAlbumContentsDTO = append(postAlbumContentsDTO, convertPostAlbumContentToPostAlbumContentDTO(postAlbumContents[i]))
	}
	return postAlbumContentsDTO

}