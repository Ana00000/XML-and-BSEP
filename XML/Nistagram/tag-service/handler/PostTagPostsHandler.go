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

type PostTagPostsHandler struct {
	Service *service.PostTagPostsService
}

func (handler *PostTagPostsHandler) CreatePostTagPosts(w http.ResponseWriter, r *http.Request) {
	var postTagPostsDTO dto.PostTagPostsDTO
	err := json.NewDecoder(r.Body).Decode(&postTagPostsDTO)

	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostTagPostsHandler) FindAllTagsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPost(&singlePostDTO)

	tagsForPostJson, _ := json.Marshal(tags)
	w.Write(tagsForPostJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostTagPostsHandler) FindPostIdsByTagId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagId := vars["tagID"]
	//var listIds []uuid.UUID
	var tags = handler.Service.FindAllPostIdsWithTagId(uuid.MustParse(tagId))
	fmt.Println("LISTA TAGOVA ID")
	fmt.Println(tags)

	fmt.Println("SALJE SE-->")
	fmt.Println(tags)
	tagsForPostJson, _ := json.Marshal(tags)
	w.Write(tagsForPostJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type ListId struct {
	ID uuid.UUID `json:"id"`
}

func (handler *PostTagPostsHandler) FindAllTagsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPosts(singlePostsDTO)

	tagsJson, _ := json.Marshal(convertListTagToListTagFullDTO(tags))
	w.Write(tagsJson)

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

func (handler *PostTagPostsHandler) FindAllTagsForPostsTagPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPostsTagPosts(singlePostsDTO)
	for i := 0; i < len(tags); i++ {
		fmt.Println("----------Naziv taga : " + tags[i].TagId.String())
	}
	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
