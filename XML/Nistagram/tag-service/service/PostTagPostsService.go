package service

import (
	"github.com/google/uuid"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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


func (service *PostTagPostsService) FindAllTagsForPosts(allPosts []postsModel.SinglePost) []model.PostTagPosts {
	tags := service.Repo.FindAllTagsForPosts(allPosts)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *PostTagPostsService) FindAllTagsForPost(post *postsModel.SinglePost) []model.PostTagPosts {
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



