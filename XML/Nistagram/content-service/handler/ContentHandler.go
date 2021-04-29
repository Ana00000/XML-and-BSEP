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

type ContentHandler struct {
	Service * service.ContentService
}

func (handler *ContentHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	var contentDTO dto.ContentDTO
	err := json.NewDecoder(r.Body).Decode(&contentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content := model.Content{
		ID:   uuid.UUID{},
		Path: contentDTO.Path,
		Type: contentDTO.Type,
	}

	err = handler.Service.CreateContent(&content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}