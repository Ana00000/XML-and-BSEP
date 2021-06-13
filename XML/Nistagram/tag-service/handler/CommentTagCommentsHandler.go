package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type CommentTagCommentsHandler struct {
	Service * service.CommentTagCommentsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
	TagService * service.TagService
}

//CRCOMMTAGCOMMTS9327
func (handler *CommentTagCommentsHandler) CreateCommentTagComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var commentTagCommentsDTO dto.CommentTagCommentsDTO
	err := json.NewDecoder(r.Body).Decode(&commentTagCommentsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "CRCOMMTAGCOMMTS9327",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CommentTagCommentsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentTagComments := model.CommentTagComments{
		ID:        uuid.UUID{},
		TagId:     commentTagCommentsDTO.TagId,
		CommentId: commentTagCommentsDTO.CommentId,
	}

	err = handler.Service.CreateCommentTagComments(&commentTagComments)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "CRCOMMTAGCOMMTS9327",
			"timestamp":   time.Now().String(),
		}).Error("Failed adding comment tag for comment!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentTagCommentsHandler",
		"action":   "CRCOMMTAGCOMMTS9327",
		"timestamp":   time.Now().String(),
	}).Info("Successfully added comment tag for comment!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentTagCommentsHandler) FindAllCommentTagCommentsForComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	id := vars["id"]
	var retValues []string
	commentsTagComments := handler.Service.FindAllCommentTagCommentsByCommentId(uuid.MustParse(id))
	for i := 0; i < len(commentsTagComments); i++ {
		var tagName = handler.TagService.FindTagNameById(commentsTagComments[i].TagId)
		retValues = append(retValues, tagName)
	}
	tagsJson, _ := json.Marshal(retValues)
	w.Write(tagsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}