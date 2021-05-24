package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
)

type TagHandler struct {
	Service * service.TagService
	Validator *validator.Validate
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
	if err := json.NewDecoder(r.Body).Decode(&tagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&tagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	findTag := handler.Service.FindTagByName(tagDTO.Name)
	var tag model.Tag
	var tagType model.TagType
	if tagDTO.TagType=="HASH_TAG"{
		tagType=model.HASH_TAG
	} else if tagDTO.TagType=="USER_TAG"{
		tagType=model.USER_TAG
	}

	if findTag != nil && findTag.TagType==tagType{
		tagJson, _ := json.Marshal(findTag.ID)
		w.Write(tagJson)
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")// 202
		return
	} else {
		var tagType model.TagType
		switch tagDTO.TagType {
		case "USER_TAG":
			tagType = model.USER_TAG
		case "HASH_TAG":
			tagType = model.HASH_TAG
		}

		tagId := uuid.New()
		tag = model.Tag{
			ID:      tagId,
			Name:    tagDTO.Name,
			TagType: tagType,
		}

		if err := handler.Service.CreateTag(&tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	tagIDJson, _ := json.Marshal(tag.ID)
	w.Write(tagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) FindTagForId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	tag := handler.Service.FindTagForId(uuid.MustParse(id))
	tagJson, _ := json.Marshal(tag)
	if tagJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(tagJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}