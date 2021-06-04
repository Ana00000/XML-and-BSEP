package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"net/http"
	_ "strconv"
)

type CommentICRHandler struct {
	Service *service.CommentICRService
}

func (handler *CommentICRHandler) CreateCommentICR(w http.ResponseWriter, r *http.Request) {
	var commentICRDTO dto.CommentICRDTO
	err := json.NewDecoder(r.Body).Decode(&commentICRDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	err = handler.Service.CreateCommentICR(&commentICR)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
