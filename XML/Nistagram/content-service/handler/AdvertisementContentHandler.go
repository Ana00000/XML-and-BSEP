package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"net/http"
	_ "strconv"
)

type AdvertisementContentHandler struct {
	Service *service.AdvertisementContentService
}

func (handler *AdvertisementContentHandler) CreateAdvertisementContent(w http.ResponseWriter, r *http.Request) {
	var advertisementContentDTO dto.AdvertisementContentDTO
	err := json.NewDecoder(r.Body).Decode(&advertisementContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	advertisementContent := model.AdvertisementContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: advertisementContentDTO.Path,
			Type: advertisementContentDTO.Type,
		},
		Link:            advertisementContentDTO.Link,
		AdvertisementId: advertisementContentDTO.AdvertisementId,
	}

	err = handler.Service.CreateAdvertisementContent(&advertisementContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
