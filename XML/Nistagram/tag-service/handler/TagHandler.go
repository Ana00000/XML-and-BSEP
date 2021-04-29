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

type TagHandler struct {
	Service * service.TagService
}

func (handler *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tagDTO dto.TagDTO
	err := json.NewDecoder(r.Body).Decode(&tagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tag := model.Tag{
		ID:          uuid.UUID{},
		Name: 		tagDTO.Name,
	}

	err = handler.Service.CreateTag(&tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}