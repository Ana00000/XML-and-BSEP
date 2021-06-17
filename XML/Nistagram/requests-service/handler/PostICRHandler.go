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

type PostICRHandler struct {
	Service   *service.PostICRService
	InappropriateContentRequestService   *service.InappropriateContentRequestService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *PostICRHandler) CreatePostICR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var postICRDTO dto.PostICRDTO
	if err := json.NewDecoder(r.Body).Decode(&postICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "PostICRHandler",
			"action":    "CREPOSTICR2544",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to PostICRDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&postICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "PostICRHandler",
			"action":    "CREPOSTICR2544",
			"timestamp": time.Now().String(),
		}).Error("PostICRDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	postICR := model.PostICR{
		InappropriateContentRequest: model.InappropriateContentRequest{
			ID:     uuid.UUID{},
			Note:   postICRDTO.Note,
			UserId: postICRDTO.UserId,
		},
		PostId: postICRDTO.PostId,
	}

	if err := handler.Service.CreatePostICR(&postICR); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "PostICRHandler",
			"action":    "CREPOSTICR2544",
			"timestamp": time.Now().String(),
		}).Error("Failed creating post inappropriate content request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.InappropriateContentRequestService.CreateICR(&postICR.InappropriateContentRequest); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "PostICRHandler",
			"action":    "CREPOSTICR2544",
			"timestamp": time.Now().String(),
		}).Error("Failed creating inappropriate content request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "PostICRHandler",
		"action":    "CREPOSTICR2544",
		"timestamp": time.Now().String(),
	}).Info("Successfully created post inappropriate content request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
