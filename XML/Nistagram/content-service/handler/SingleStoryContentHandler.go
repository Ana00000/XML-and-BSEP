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

func (handler *SingleStoryContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myStoryFile")
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

	pathStoryGlobal = tempFile.Name()[6:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}
