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

type TagHandler struct {
	Service * service.TagService
}

type ReturnValueString struct {
	ReturnValue string `json:"return_value"`
}

func (handler *TagHandler) FindTagNameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var tagName = handler.Service.FindTagNameById(uuid.MustParse(id))

	returnValue := ReturnValueString{ReturnValue: tagName}

	returnValueJson, _ := json.Marshal(returnValue)
	w.Write(returnValueJson)


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tagDTO dto.TagDTO
	err := json.NewDecoder(r.Body).Decode(&tagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tag := model.Tag{
		ID:          uuid.UUID{},
		Name: 		tagDTO.Name,
	}

	err = handler.Service.CreateTag(&tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}