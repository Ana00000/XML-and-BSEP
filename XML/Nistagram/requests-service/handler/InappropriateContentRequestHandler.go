package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type InappropriateContentRequestHandler struct {
	Service   *service.InappropriateContentRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *InappropriateContentRequestHandler) CreateInappropriateContentRequest(w http.ResponseWriter, r *http.Request) {
	var inappropriateContentRequestDTO dto.InappropriateContentRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&inappropriateContentRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "InappropriateContentRequestHandler",
			"action":    "CRINAPPROPCONTREQ4255",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to InappropriateContentRequestDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&inappropriateContentRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "InappropriateContentRequestHandler",
			"action":    "CRINAPPROPCONTREQ4255",
			"timestamp": time.Now().String(),
		}).Error("InappropriateContentRequestDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	inappropriateContentRequest := model.InappropriateContentRequest{
		ID:     uuid.UUID{},
		Note:   inappropriateContentRequestDTO.Note,
		UserId: inappropriateContentRequestDTO.UserId,
	}

	if err := handler.Service.CreateInappropriateContentRequest(&inappropriateContentRequest); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "InappropriateContentRequestHandler",
			"action":    "CRINAPPROPCONTREQ4255",
			"timestamp": time.Now().String(),
		}).Error("Failed creating inappropriate content request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "InappropriateContentRequestHandler",
		"action":    "CRINAPPROPCONTREQ4255",
		"timestamp": time.Now().String(),
	}).Info("Successfully created inappropriate content request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
