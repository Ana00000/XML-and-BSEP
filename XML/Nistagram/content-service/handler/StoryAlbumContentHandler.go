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

type StoryAlbumContentHandler struct {
	Service * service.StoryAlbumContentService
	ContentService * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

var pathStoryAlbumGlobal = ""

func (handler *StoryAlbumContentHandler) CreateStoryAlbumContent(w http.ResponseWriter, r *http.Request) {
	var storyAlbumContentDTO dto.StoryAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "CRSTALCOX866",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch storyAlbumContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	storyAlbumContent := model.StoryAlbumContent{
		Content: model.Content{
			ID:   id,
			Path: pathStoryAlbumGlobal,
			Type: contentType,
		},
		StoryAlbumId: storyAlbumContentDTO.StoryAlbumId,
	}

	err = handler.Service.CreateStoryAlbumContent(&storyAlbumContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "StoryAlbumContentHandler",
		"action":   "CRSTALCOX866",
		"timestamp":   time.Now().String(),
	}).Error("Failed creating story album content!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&storyAlbumContent.Content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "CRSTALCOX866",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumContentHandler",
		"action":   "CRSTALCOX866",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story album content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myStoryAlbumFile")
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "StoryAlbumContentHandler",
		"action":   "UPQ799",
		"timestamp":   time.Now().String(),
	}).Error("Failed to find the file!")
		fmt.Println(err)
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"),  "*" + hand.Filename)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "UPQ799",
			"timestamp":   time.Now().String(),
		}).Error("Failed to create temporary file!")
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "UPQ799",
			"timestamp":   time.Now().String(),
		}).Error("Failed to read from file!")
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	pathStoryAlbumGlobal = tempFile.Name()[20:len(tempFile.Name())]
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumContentHandler",
		"action":   "UPQ799",
		"timestamp":   time.Now().String(),
	}).Info("Successfully uploaded the media!")
	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func (handler *StoryAlbumContentHandler) FindAllContentsForStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO []dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "FIALCOFOSTALJ710",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStoryAlbums = convertListStoryAlbumContentToListStoryAlbumContentDTO(handler.Service.FindAllContentsForStoryAlbums(storyAlbumFullDTO))

	contentsForStoryAlbumsJson, _ := json.Marshal(contentsForStoryAlbums)
	w.Write(contentsForStoryAlbumsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumContentHandler",
		"action":   "FIALCOFOSTALJ710",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for story albums!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumContentHandler) FindAllContentsForStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryAlbumContentHandler",
			"action":   "FIALCOFOSTALS400",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryAlbumFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStoryAlbums = convertListStoryAlbumContentToListStoryAlbumContentDTO(handler.Service.FindAllContentsForStoryAlbum(&storyAlbumFullDTO))

	contentsForStoryAlbumsJson, _ := json.Marshal(contentsForStoryAlbums)
	w.Write(contentsForStoryAlbumsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryAlbumContentHandler",
		"action":   "FIALCOFOSTALS400",
		"timestamp":   time.Now().String(),
	}).Info("Successfully found contents for story album!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertStoryAlbumContentToStoryAlbumContentDTO(storyAlbumContent model.StoryAlbumContent) dto.StoryAlbumContentFullDTO{
	storyAlbumContentType :=""
	if storyAlbumContent.Type==model.PICTURE{
		storyAlbumContentType="PICTURE"
	}else if storyAlbumContent.Type==model.VIDEO{
		storyAlbumContentType="VIDEO"
	}
	var storyAlbumContentDTO = dto.StoryAlbumContentFullDTO{
		ID:           storyAlbumContent.ID,
		Path:         storyAlbumContent.Path,
		Type:         storyAlbumContentType,
		StoryAlbumId: storyAlbumContent.StoryAlbumId,
	}
	return storyAlbumContentDTO
}

func convertListStoryAlbumContentToListStoryAlbumContentDTO(storyAlbumContents []model.StoryAlbumContent) []dto.StoryAlbumContentFullDTO{
	var storyAlbumContentsDTO []dto.StoryAlbumContentFullDTO
	for i := 0; i < len(storyAlbumContents); i++ {
		storyAlbumContentsDTO=append(storyAlbumContentsDTO, convertStoryAlbumContentToStoryAlbumContentDTO(storyAlbumContents[i]))
	}
	return storyAlbumContentsDTO
}