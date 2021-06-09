package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type TagHandler struct {
	Service * service.TagService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

type ReturnValueString struct {
	ReturnValue string `json:"return_value"`
}
//FIDTGNMBYID0921
func (handler *TagHandler) FindTagNameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//fmt.Println("Dobijeni ID : "+id)
	var tagName = handler.Service.FindTagNameById(uuid.MustParse(id))
	//fmt.Println("Dobijeni name : "+tagName)
	returnValue := ReturnValueString{ReturnValue: tagName}

	returnValueJson, _ := json.Marshal(returnValue)
	w.Write(returnValueJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "TagHandler",
		"action":   "FIDTGNMBYID0921",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tag name by id!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
//CRTG7821
func (handler *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tagDTO dto.TagDTO
	if err := json.NewDecoder(r.Body).Decode(&tagDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagHandler",
			"action":   "CRTG7821",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to TagDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&tagDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagHandler",
			"action":   "CRTG7821",
			"timestamp":   time.Now().String(),
		}).Error("TagDTO fields doesn't in valid format!")
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
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "TagHandler",
			"action":   "CRTG7821",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded tag with same name!")
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
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "TagHandler",
				"action":   "CRTG7821",
				"timestamp":   time.Now().String(),
			}).Error("Failed creating tag!")
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	tagIDJson, _ := json.Marshal(tag.ID)
	w.Write(tagIDJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "TagHandler",
		"action":   "CRTG7821",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created tag!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
//FIDTGFORID9180
func (handler *TagHandler) FindTagForId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	tag := handler.Service.FindTagForId(uuid.MustParse(id))
	tagJson, _ := json.Marshal(tag)
	if tagJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(tagJson)
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "TagHandler",
			"action":   "FIDTGFORID9180",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded tag for id!")
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "TagHandler",
		"action":   "FIDTGFORID9180",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding tag for id!")
	w.WriteHeader(http.StatusBadRequest)
}
//FIDTGBYNM913
func (handler *TagHandler) FindTagByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	tag := handler.Service.FindTagByName(name)
	tagType :=""
	if tag.TagType==model.USER_TAG{
		tagType="USER_TAG"
	} else if tag.TagType==model.HASH_TAG{
		tagType="HASH_TAG"
	}
	tagDTO := dto.TagFullDTO{
		ID:      tag.ID,
		Name:    tag.Name,
		TagType: tagType,
	}
	tagJson, _ := json.Marshal(tagDTO)
	if tagJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(tagJson)
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "TagHandler",
			"action":   "FIDTGBYNM913",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded tag by name!")
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "TagHandler",
		"action":   "FIDTGBYNM913",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding tag by name!")
	w.WriteHeader(http.StatusBadRequest)
}