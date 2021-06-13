package handler

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"os"
	"time"
)

type CommentHandler struct {
	Service   *service.CommentService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-comment-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var commentDTO dto.CommentDTO
	if err := json.NewDecoder(r.Body).Decode(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CommentDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("CommentDTO fields aren't in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}
	comment := model.Comment{
		ID:           uuid.UUID{},
		CreationDate: time.Now(),
		UserID:       commentDTO.UserID,
		PostID:       commentDTO.PostID,
		Text:         commentDTO.Text,
	}

	if err := handler.Service.CreateComment(&comment); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating comment!")
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentHandler",
		"action":   "CRCOM571",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created comment!")
	
	commentIDJson, _ := json.Marshal(comment.ID)
	w.Write(commentIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) FindAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-comments-for-post-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllCommentsForPost(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all comments for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FACFP572",
		"timestamp":   time.Now().String(),
	}).Error("Comments for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *CommentHandler) FindAllUserComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllUserComments(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FAUCM573",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all user comments!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FAUCM573",
		"timestamp":   time.Now().String(),
	}).Error("All user comments not found!")
	w.WriteHeader(http.StatusBadRequest)
}
