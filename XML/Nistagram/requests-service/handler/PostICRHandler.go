package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type PostICRHandler struct {
	Service * service.PostICRService
}

func (handler *PostICRHandler) CreatePostICR(w http.ResponseWriter, r *http.Request) {
	var postICRDTO dto.PostICRDTO
	err := json.NewDecoder(r.Body).Decode(&postICRDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postICR := model.PostICR{
		InappropriateContentRequest : model.InappropriateContentRequest{
			ID:          uuid.UUID{},
			Note: 		 postICRDTO.Note,
			UserId:      postICRDTO.UserId,
		},
		PostId:      postICRDTO.PostId,
	}

	err = handler.Service.CreatePostICR(&postICR)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

