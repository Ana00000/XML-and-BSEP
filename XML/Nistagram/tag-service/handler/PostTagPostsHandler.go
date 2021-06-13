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

type PostTagPostsHandler struct {
	Service * service.PostTagPostsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CRPOTGPSTS532
func (handler *PostTagPostsHandler) CreatePostTagPosts(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "CRPOTGPSTS532",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-post-tag-posts-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "CRPOTGPSTS532",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var postTagPostsDTO dto.PostTagPostsDTO
	err := json.NewDecoder(r.Body).Decode(&postTagPostsDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "CRPOTGPSTS532",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostTagPostsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postTagPosts := model.PostTagPosts{
		ID:     uuid.UUID{},
		TagId:  postTagPostsDTO.TagId,
		PostId: postTagPostsDTO.PostId,
	}

	err = handler.Service.CreatePostTagPosts(&postTagPosts)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "CRPOTGPSTS532",
			"timestamp":   time.Now().String(),
		}).Error("Failed adding post tag for post!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostTagPostsHandler",
		"action":   "CRPOTGPSTS532",
		"timestamp":   time.Now().String(),
	}).Info("Successfully added post tag for post!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALTGSFORPST9128
func (handler *PostTagPostsHandler) FindAllTagsForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "FIDALTGSFORPST9128",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPost(&singlePostDTO)

	tagsForPostJson, _ := json.Marshal(tags)
	w.Write(tagsForPostJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostTagPostsHandler",
		"action":   "FIDALTGSFORPST9128",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tags for post!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDPSTIDSBYTGID9851
func (handler *PostTagPostsHandler) FindPostIdsByTagId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	tagId := vars["tagID"]
	//var listIds []uuid.UUID
	var tags = handler.Service.FindAllPostIdsWithTagId(uuid.MustParse(tagId))
	tagsForPostJson, _ := json.Marshal(tags)
	//fmt.Println("ONO STO IDE NA BEK ----->")
	//fmt.Println(string(tagsForPostJson))
	w.Write(tagsForPostJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostTagPostsHandler",
		"action":   "FIDPSTIDSBYTGID9851",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded post ids with tag id!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type ListId struct {
	ID uuid.UUID `json:"id"`
}

//FIDALTGSFORPSTS9882
func (handler *PostTagPostsHandler) FindAllTagsForPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "FIDALTGSFORPSTS9882",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags =  handler.Service.FindAllTagsForPosts(singlePostsDTO)

	tagsJson, _ := json.Marshal(convertListTagToListTagFullDTO(tags))
	w.Write(tagsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostTagPostsHandler",
		"action":   "FIDALTGSFORPSTS9882",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tags for posts!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertTagToTagFullDTO(tag model.Tag) dto.TagFullDTO {
	tagType := ""
	if tag.TagType == model.HASH_TAG {
		tagType = "HASH_TAG"
	} else if tag.TagType == model.USER_TAG {
		tagType = "USER_TAG"
	}
	var tagDTO = dto.TagFullDTO{
		ID:      tag.ID,
		Name:    tag.Name,
		TagType: tagType,
	}
	return tagDTO
}

func convertListTagToListTagFullDTO(tag []model.Tag) []dto.TagFullDTO {
	var listTagDTO []dto.TagFullDTO
	for i := 0; i < len(tag); i++ {
		listTagDTO = append(listTagDTO, convertTagToTagFullDTO(tag[i]))
	}
	return listTagDTO
}

//FIDALTGSFORPSTSTGPSTS9
func (handler *PostTagPostsHandler) FindAllTagsForPostsTagPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostTagPostsHandler",
			"action":   "FIDALTGSFORPSTSTGPSTS9",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to list SinglePostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tags = handler.Service.FindAllTagsForPostsTagPosts(singlePostsDTO)
	/*for i := 0; i < len(tags); i++ {
		fmt.Println("----------Naziv taga : "+tags[i].TagId.String())
	}*/
	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostTagPostsHandler",
		"action":   "FIDALTGSFORPSTS9882",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded tags for posts tag for posts!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
