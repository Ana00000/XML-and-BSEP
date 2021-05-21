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

type PostAlbumContentHandler struct {
	Service * service.PostAlbumContentService
	ContentService * service.ContentService
}

var pathPostAlbumGlobal = ""

func (handler *PostAlbumContentHandler) CreatePostAlbumContent(w http.ResponseWriter, r *http.Request) {
	var postAlbumContentDTO dto.PostAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumContentDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&postAlbumContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostAlbumFile")
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

	pathPostAlbumGlobal = tempFile.Name()[6:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}
