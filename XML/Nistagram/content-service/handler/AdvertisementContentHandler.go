package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type AdvertisementContentHandler struct {
	Service *service.AdvertisementContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *AdvertisementContentHandler) CreateAdvertisementContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var advertisementContentDTO dto.AdvertisementContentDTO
	err := json.NewDecoder(r.Body).Decode(&advertisementContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdvertisementContentHandler",
			"action":   "CRADCOO454",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to AdvertisementContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch advertisementContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	advertisementContent := model.AdvertisementContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: advertisementContentDTO.Path,
			Type: contentType,
		},
		Link:            advertisementContentDTO.Link,
		AdvertisementId: advertisementContentDTO.AdvertisementId,
	}

	err = handler.Service.CreateAdvertisementContent(&advertisementContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdvertisementContentHandler",
			"action":   "CRADCOO454",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating advertisement content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "AdvertisementContentHandler",
		"action":   "CRADCOO454",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created advertisement content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
