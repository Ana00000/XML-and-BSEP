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
	_ "strconv"
)

type SinglePostContentHandler struct {
	Service * service.SinglePostContentService
	ContentService * service.ContentService
}

var pathPostGlobal = ""

func (handler *SinglePostContentHandler) CreateSinglePostContent(w http.ResponseWriter, r *http.Request) {
	var singlePostContentDTO dto.SinglePostContentDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch singlePostContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	singlePostContent := model.SinglePostContent{
		Content: model.Content{
			ID:   id,
			Path: pathPostGlobal,
			Type: contentType,
		},
		SinglePostId: singlePostContentDTO.SinglePostId,
	}

	err = handler.Service.CreateSinglePostContent(&singlePostContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&singlePostContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	pathPostGlobal = ""

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostFile")
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


	pathPostGlobal = tempFile.Name()[6:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}