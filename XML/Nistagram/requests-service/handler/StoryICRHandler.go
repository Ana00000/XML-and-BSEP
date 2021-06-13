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

type StoryICRHandler struct {
	Service   *service.StoryICRService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *StoryICRHandler) CreateStoryICR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var storyICRDTO dto.StoryICRDTO

	if err := json.NewDecoder(r.Body).Decode(&storyICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryICRHandler",
			"action":    "CRESTORYICR3443",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to StoryICRDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&storyICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryICRHandler",
			"action":    "CRESTORYICR3443",
			"timestamp": time.Now().String(),
		}).Error("StoryICRDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	storyICR := model.StoryICR{
		InappropriateContentRequest: model.InappropriateContentRequest{
			ID:     uuid.UUID{},
			Note:   storyICRDTO.Note,
			UserId: storyICRDTO.UserId,
		},
		StoryId: storyICRDTO.StoryId,
	}


	if err := handler.Service.CreateStoryICR(&storyICR); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryICRHandler",
			"action":    "CRESTORYICR3443",
			"timestamp": time.Now().String(),
		}).Error("Failed creating story inappropriate content request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "StoryICRHandler",
		"action":    "CRESTORYICR3443",
		"timestamp": time.Now().String(),
	}).Info("Successfully created story inappropriate content request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
