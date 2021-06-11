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

type VerificationRequestHandler struct {
	Service   *service.VerificationRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *VerificationRequestHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request) {
	var verificationRequestDTO dto.VerificationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&verificationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to VerificationRequestDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}
	if err := handler.Validator.Struct(&verificationRequestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("VerificationRequestDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	verificationRequest := model.VerificationRequest{
		ID:                     uuid.UUID{},
		FirstName:              verificationRequestDTO.FirstName,
		LastName:               verificationRequestDTO.LastName,
		OfficialDocumentPath:   verificationRequestDTO.OfficialDocumentPath,
		RegisteredUserCategory: verificationRequestDTO.RegisteredUserCategory,
	}

	if err := handler.Service.CreateVerificationRequest(&verificationRequest); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "VerificationRequestHandler",
			"action":    "CREVERIFREQ6631",
			"timestamp": time.Now().String(),
		}).Error("Failed creating verification request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "VerificationRequestHandler",
		"action":    "CREVERIFREQ6631",
		"timestamp": time.Now().String(),
	}).Info("Successfully created verification request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
