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

type SinglePostContentHandler struct {
	Service        *service.SinglePostContentService
	ContentService *service.ContentService
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

func (handler *SinglePostContentHandler) FindAllContentsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTOs []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTOs)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contents = convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(handler.Service.FindAllContentsForPosts(singlePostsDTOs))

	pathJson, _ := json.Marshal(contents)
	w.Write(pathJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) FindAllContentsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contents = convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(handler.Service.FindAllContentsForPost(&singlePostDTO))

	pathJson, _ := json.Marshal(contents)
	w.Write(pathJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SinglePostContentHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, hand, err := request.FormFile("myPostFile")
	if err != nil {
		fmt.Println("Greska prvi if")
		fmt.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(os.Getenv("BASE_URL"), "*"+hand.Filename)
	if err != nil {
		fmt.Println("Greska drugi if")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Greska treci if")
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	pathPostGlobal = tempFile.Name()[20:len(tempFile.Name())]

	pathJson, _ := json.Marshal(tempFile.Name())
	writer.Write(pathJson)
}

func convertListSinglePostContentsToListSinglePostContentsForSinglePostDTO(singlePostContents []model.SinglePostContent) []dto.SinglePostContentForSinglePostDTO {
	var listSinglePostContentsDTO []dto.SinglePostContentForSinglePostDTO
	for i := 0; i < len(singlePostContents); i++ {
		listSinglePostContentsDTO = append(listSinglePostContentsDTO, convertSinglePostContentToSinglePostContentForSinglePostDTO(singlePostContents[i]))
	}
	return listSinglePostContentsDTO
}

func convertSinglePostContentToSinglePostContentForSinglePostDTO(singlePostContent model.SinglePostContent) dto.SinglePostContentForSinglePostDTO {
	contentType := ""
	if singlePostContent.Type == model.VIDEO {
		contentType = "VIDEO"
	} else if singlePostContent.Type == model.PICTURE {
		contentType = "PICTURE"
	}

	var singlePostContentDTO = dto.SinglePostContentForSinglePostDTO{
		ID:           singlePostContent.ID,
		Path:         singlePostContent.Path,
		Type:         contentType,
		SinglePostId: singlePostContent.SinglePostId,
	}
	return singlePostContentDTO
}
