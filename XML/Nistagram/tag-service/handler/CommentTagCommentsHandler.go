package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	"os"
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
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "CRCOMMTAGCOMMTS9327",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-comment-tag-comments-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "CRCOMMTAGCOMMTS9327",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

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

func Request(url string, token string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	tokenString := "Bearer "+token
	req.Header.Set("Authorization", tokenString)
	resp, err := http.DefaultClient.Do(req)
	return resp
}

//FIDALCOMMTGCOMMTSFORCOMM9027
func (handler *CommentTagCommentsHandler) FindAllCommentTagCommentsForComment(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "FIDALCOMMTGCOMMTSFORCOMM9027",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-comment-tag-comments-for-comment-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentTagCommentsHandler",
			"action":   "FIDALCOMMTGCOMMTSFORCOMM9027",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

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

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentTagCommentsHandler",
		"action":   "FIDALCOMMTGCOMMTSFORCOMM9027",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded comment tag comments for comment!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}