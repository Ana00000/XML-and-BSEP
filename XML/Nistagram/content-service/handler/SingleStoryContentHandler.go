package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
)

type SingleStoryContentHandler struct {
	Service * service.SingleStoryContentService
	ContentService * service.ContentService
}

var pathStoryGlobal = ""

func (handler *SingleStoryContentHandler) CreateSingleStoryContent(w http.ResponseWriter, r *http.Request) {
	var singleStoryContentDTO dto.SingleStoryContentDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryContentDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&singleStoryContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	pathStoryGlobal = ""

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) FindAllContentsForStories(w http.ResponseWriter, r *http.Request) {
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStories = convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(handler.Service.FindAllContentsForStories(singleStoriesDTO))

	contentsForStoriesJson, _ := json.Marshal(contentsForStories)
	w.Write(contentsForStoriesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) FindAllContentsForStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForStories = convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(handler.Service.FindAllContentsForStory(&singleStoryDTO))
	contentsForStoriesJson, _ := json.Marshal(contentsForStories)

	w.Write(contentsForStoriesJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myStoryFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"),  "*" + hand.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	pathStoryGlobal = tempFile.Name()[20:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func convertSingleStoryContentToSingleStoryContentForSingleStoryDTO(singleStoryContent model.SingleStoryContent) dto.SingleStoryContentForSingleStoryDTO{
	contentType:= ""
	if singleStoryContent.Type==model.VIDEO{
		contentType="VIDEO"
	} else if singleStoryContent.Type==model.PICTURE{
		contentType="PICTURE"
	}
	var singleStoryContentForSingleStoryDTO = dto.SingleStoryContentForSingleStoryDTO{
		ID:            singleStoryContent.ID,
		Path:          singleStoryContent.Path,
		Type:          contentType,
		SingleStoryId: singleStoryContent.SingleStoryId,
	}
	return singleStoryContentForSingleStoryDTO
}

func convertListSingleStoriesContentToSingleStoriesContentForSingleStoryDTO(singleStoryContents []model.SingleStoryContent) []dto.SingleStoryContentForSingleStoryDTO{
	var listSingleStoryContentForSingleStoryDTO []dto.SingleStoryContentForSingleStoryDTO
	for i := 0; i < len(singleStoryContents); i++ {
		listSingleStoryContentForSingleStoryDTO = append(listSingleStoryContentForSingleStoryDTO, convertSingleStoryContentToSingleStoryContentForSingleStoryDTO(singleStoryContents[i]))
	}
	return listSingleStoryContentForSingleStoryDTO
}