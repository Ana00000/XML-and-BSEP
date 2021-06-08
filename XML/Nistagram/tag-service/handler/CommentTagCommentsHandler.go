package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
)

type CommentTagCommentsHandler struct {
	Service * service.CommentTagCommentsService
	TagService * service.TagService
}

func (handler *CommentTagCommentsHandler) CreateCommentTagComments(w http.ResponseWriter, r *http.Request) {
	var commentTagCommentsDTO dto.CommentTagCommentsDTO
	err := json.NewDecoder(r.Body).Decode(&commentTagCommentsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentTagComments := model.CommentTagComments{
		ID:          uuid.UUID{},
		TagId: 		commentTagCommentsDTO.TagId,
		CommentId: commentTagCommentsDTO.CommentId,
	}

	err = handler.Service.CreateCommentTagComments(&commentTagComments)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentTagCommentsHandler) FindAllCommentTagCommentsForComment(w http.ResponseWriter, r *http.Request) {
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