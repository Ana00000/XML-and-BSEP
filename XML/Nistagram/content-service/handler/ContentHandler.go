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

type ContentHandler struct {
	Service * service.ContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ContentHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var contentDTO dto.ContentDTO
	err := json.NewDecoder(r.Body).Decode(&contentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentHandler",
			"action":   "CRCOU658",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch contentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	content := model.Content{
		ID:   id,
		Path: contentDTO.Path,
		Type: contentType,
	}

	err = handler.Service.CreateContent(&content)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentHandler",
			"action":   "CRCOU658",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ContentHandler",
		"action":   "CRCOU658",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
