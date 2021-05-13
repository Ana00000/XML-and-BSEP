package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type SingleStoryContentHandler struct {
	Service * service.SingleStoryContentService
}

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

	singleStoryContent := model.SingleStoryContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: singleStoryContentDTO.Path,
			Type: contentType,
		},
		SingleStoryId: singleStoryContentDTO.SingleStoryId,
	}

	err = handler.Service.CreateSingleStoryContent(&singleStoryContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
