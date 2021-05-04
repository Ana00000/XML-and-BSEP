package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type PostTagHandler struct {
	Service * service.PostTagService
}

func (handler *PostTagHandler) CreatePostTag(w http.ResponseWriter, r *http.Request) {
	var postTagDTO dto.PostTagDTO
	err := json.NewDecoder(r.Body).Decode(&postTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postTag := model.PostTag{
		Tag: model.Tag{
			ID:   uuid.UUID{},
			Name: postTagDTO.Name,
		},
	}

	err = handler.Service.CreatePostTag(&postTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
