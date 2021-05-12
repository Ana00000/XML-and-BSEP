package service

import (
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
