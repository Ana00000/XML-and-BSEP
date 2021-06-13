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

type CommentICRHandler struct {
	Service   *service.CommentICRService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *CommentICRHandler) CreateCommentICR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var commentICRDTO dto.CommentICRDTO
	if err := json.NewDecoder(r.Body).Decode(&commentICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "CommentICRHandler",
			"action":    "CRCOMICR9998",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to CommentICRDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&commentICRDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "CommentICRHandler",
			"action":    "CRCOMICR9998",
			"timestamp": time.Now().String(),
		}).Error("CommentICRDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	commentICR := model.CommentICR{
		InappropriateContentRequest: model.InappropriateContentRequest{
			ID:     uuid.UUID{},
			Note:   commentICRDTO.Note,
			UserId: commentICRDTO.UserId,
		},
		CommentId: commentICRDTO.CommentId,
	}

	if err := handler.Service.CreateCommentICR(&commentICR); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "CommentICRHandler",
			"action":    "CRCOMICR9998",
			"timestamp": time.Now().String(),
		}).Error("Failed creating comment inappropriate content request!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "CommentICRHandler",
		"action":    "CRCOMICR9998",
		"timestamp": time.Now().String(),
	}).Info("Successfully created comment inappropriate content request!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
