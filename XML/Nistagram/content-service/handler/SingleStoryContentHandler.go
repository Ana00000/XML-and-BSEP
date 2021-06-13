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

type SingleStoryContentHandler struct {
	Service * service.SingleStoryContentService
	ContentService * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

var pathStoryGlobal = ""

func (handler *SingleStoryContentHandler) CreateSingleStoryContent(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "CRSISTCOA197",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-single-story-content-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "CRSISTCOA197",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singleStoryContentDTO dto.SingleStoryContentDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "CRSISTCOA197",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch singleStoryContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	singleStoryContent := model.SingleStoryContent{
		Content: model.Content{
			ID:   id,
			Path: pathStoryGlobal,
			Type: contentType,
		},
		SingleStoryId: singleStoryContentDTO.SingleStoryId,
	}

	err = handler.Service.CreateSingleStoryContent(&singleStoryContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "CRSISTCOA197",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating single story content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.ContentService.CreateContent(&singleStoryContent.Content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "CRSISTCOA197",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	pathStoryGlobal = ""
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryContentHandler",
		"action":   "CRSISTCOA197",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created single story content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) FindAllContentsForStories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "FIALCOFOSTF439",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoriesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStories = convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(handler.Service.FindAllContentsForStories(singleStoriesDTO))

	contentsForStoriesJson, _ := json.Marshal(contentsForStories)
	w.Write(contentsForStoriesJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryContentHandler",
		"action":   "FIALCOFOSTF439",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for stories!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) FindAllContentsForStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "FIALCOFOSTV496",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStories = convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(handler.Service.FindAllContentsForStory(&singleStoryDTO))
	contentsForStoriesJson, _ := json.Marshal(contentsForStories)

	w.Write(contentsForStoriesJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryContentHandler",
		"action":   "FIALCOFOSTV496",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for story!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("X-XSS-Protection", "1; mode=block")
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myStoryFile")
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "UPM253",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find the file!")
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"), "*"+hand.Filename)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "UPM253",
			"timestamp":   time.Now().String(),
		}).Error("Failed to create temporary file!")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryContentHandler",
			"action":   "UPM253",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		return
	}
	tempFile.Write(fileBytes)

	pathStoryGlobal = tempFile.Name()[20:len(tempFile.Name())]
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryContentHandler",
		"action":   "UPM253",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func convertSingleStoryContentToSingleStoryContentForSingleStoryDTO(singleStoryContent model.SingleStoryContent) dto.SingleStoryContentForSingleStoryDTO {
	contentType := ""
	if singleStoryContent.Type == model.VIDEO {
		contentType = "VIDEO"
	} else if singleStoryContent.Type == model.PICTURE {
		contentType = "PICTURE"
	}
	var singleStoryContentForSingleStoryDTO = dto.SingleStoryContentForSingleStoryDTO{
		ID:            singleStoryContent.ID,
		Path:          singleStoryContent.Path,
		Type:          contentType,
		SingleStoryId: singleStoryContent.SingleStoryId,
	}
	return singleStoryContentForSingleStoryDTO
}

func convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(singleStoryContents []model.SingleStoryContent) []dto.SingleStoryContentForSingleStoryDTO {
	var listSingleStoryContentForSingleStoryDTO []dto.SingleStoryContentForSingleStoryDTO
	for i := 0; i < len(singleStoryContents); i++ {
		listSingleStoryContentForSingleStoryDTO = append(listSingleStoryContentForSingleStoryDTO, convertSingleStoryContentToSingleStoryContentForSingleStoryDTO(singleStoryContents[i]))
	}
	return listSingleStoryContentForSingleStoryDTO
}

func (handler *SingleStoryContentHandler) FindSingleStoryContentForStoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	singleStoryContent := handler.Service.FindSingleStoryContentForStoryId(uuid.MustParse(id))
	singleStoryContentJson, _ := json.Marshal(singleStoryContent)
	if singleStoryContentJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryContentHandler",
			"action":   "FISISTCOFOSTH439",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found single story content for story id!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(singleStoryContentJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "SingleStoryContentHandler",
		"action":   "FISISTCOFOSTH439",
		"timestamp":   time.Now().String(),
	}).Error("Single story content for story id wasn't found!")
	w.WriteHeader(http.StatusBadRequest)
}
