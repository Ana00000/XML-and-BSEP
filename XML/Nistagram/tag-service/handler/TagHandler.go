package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
)

type TagHandler struct {
	Service * service.TagService
}

func (handler *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tagDTO dto.TagDTO
	if err := json.NewDecoder(r.Body).Decode(&tagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	findTag := handler.Service.FindTagByName(tagDTO.Name)
	var tag model.Tag

	if findTag != nil {
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	} else {
		var tagType model.TagType
		switch tagDTO.TagType {
		case "USER_TAG":
			tagType = model.USER_TAG
		case "HASH_TAG":
			tagType = model.HASH_TAG
		}

		tagId := uuid.New()
		tag = model.Tag{
			ID:      tagId,
			Name:    tagDTO.Name,
			TagType: tagType,
		}

		if err := handler.Service.CreateTag(&tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}