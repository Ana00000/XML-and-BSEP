package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
)

type UserTagHandler struct {
	Service *service.UserTagService
	TagService *service.TagService
}

func (handler *UserTagHandler) CreateUserTag(w http.ResponseWriter, r *http.Request) {
	var userTagDTO dto.UserTagDTO
	if err := json.NewDecoder(r.Body).Decode(&userTagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var findTag = handler.TagService.FindTagByName(userTagDTO.Name)
	var userTag model.UserTag

	if findTag != nil && userTagDTO.TagType == "USER_TAG" {
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	} else {
		id := uuid.New()
		userTag = model.UserTag{
			Tag: model.Tag{
				ID: id,
				Name: userTagDTO.Name,
				TagType: model.USER_TAG,
			},
			UserId: userTagDTO.UserId,
		}

		if err := handler.Service.CreateUserTag(&userTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}

		if err := handler.TagService.CreateTag(&userTag.Tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	userTagIDJson, _ := json.Marshal(userTag.ID)
	w.Write(userTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
