package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	_ "strconv"
)

type StoryAlbumContentHandler struct {
	Service * service.StoryAlbumContentService
	ContentService * service.ContentService
}

var pathStoryAlbumGlobal = ""

func (handler *StoryAlbumContentHandler) CreateStoryAlbumContent(w http.ResponseWriter, r *http.Request) {
	var storyAlbumContentDTO dto.StoryAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumContentDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&storyAlbumContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	pathStoryAlbumGlobal = ""

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myStoryAlbumFile")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("Media",  "*" + hand.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	pathStoryAlbumGlobal = tempFile.Name()

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}
