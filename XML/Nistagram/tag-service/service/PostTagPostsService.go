package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type PostTagPostsService struct {
	Repo * repository.PostTagPostsRepository
}

func (service * PostTagPostsService) CreatePostTagPosts(postTagPosts *model.PostTagPosts) error {
	err := service.Repo.CreatePostTagPosts(postTagPosts)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostTagPostsService) FindAllTagsForPostsTagPosts(allPosts []dto.SinglePostDTO) []model.PostTagPosts {
	tags := service.Repo.FindAllTagsForPostsTagPosts(allPosts)
	if tags != nil {
		return tags
	}
	return nil
}


func (service *PostTagPostsService) FindAllTagsForPosts(allPosts []dto.SinglePostDTO) []model.Tag {
	tags := service.Repo.FindAllTagsForPosts(allPosts)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *PostTagPostsService) FindAllTagsForPost(post *dto.SinglePostDTO) []model.PostTagPosts {
	tags := service.Repo.FindAllTagsForPost(post)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *PostTagPostsService) FindAllPostIdsWithTagId(tagId uuid.UUID) []uuid.UUID {
	postIds := service.Repo.FindAllPostIdsWithTagId(tagId)
	if postIds != nil {
		return postIds
	}
	return nil
}



