package service

import (
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


func (service *PostTagPostsService) FindAllTagsForPosts(allPosts []postsModel.Post) []model.PostTagPosts {
	tags := service.Repo.FindAllTagsForPosts(allPosts)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *PostTagPostsService) FindAllTagsForPost(post *postsModel.Post) []model.PostTagPosts {
	tags := service.Repo.FindAllTagsForPost(post)
	if tags != nil {
		return tags
	}
	return nil
}

