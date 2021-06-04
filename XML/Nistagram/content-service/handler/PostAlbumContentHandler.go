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

type PostAlbumContentHandler struct {
	Service        *service.PostAlbumContentService
	ContentService *service.ContentService
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

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"), "*"+hand.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	pathPostAlbumGlobal = tempFile.Name()[20:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func (handler *PostAlbumContentHandler) FindAllContentsForPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO []dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForPostAlbums = convertListPostAlbumContentToListPostAlbumContentDTO(handler.Service.FindAllContentsForPostAlbums(postAlbumFullDTO))

	contentsForPostAlbumsJson, _ := json.Marshal(contentsForPostAlbums)
	w.Write(contentsForPostAlbumsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostAlbumContentHandler) FindAllContentsForPostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumFullDTO dto.PostAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentsForPostAlbums = convertListPostAlbumContentToListPostAlbumContentDTO(handler.Service.FindAllContentsForPostAlbum(&postAlbumFullDTO))

	contentsForPostAlbumsJson, _ := json.Marshal(contentsForPostAlbums)
	w.Write(contentsForPostAlbumsJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertPostAlbumContentToPostAlbumContentDTO(postAlbumContent model.PostAlbumContent) dto.PostAlbumContentFullDTO {
	postAlbumContentType := ""
	if postAlbumContent.Type == model.PICTURE {
		postAlbumContentType = "PICTURE"
	} else if postAlbumContent.Type == model.VIDEO {
		postAlbumContentType = "VIDEO"
	}
	var postAlbumContentDTO = dto.PostAlbumContentFullDTO{
		ID:          postAlbumContent.ID,
		Path:        postAlbumContent.Path,
		Type:        postAlbumContentType,
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
