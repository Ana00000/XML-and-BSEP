package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
)

type SinglePostContentService struct {
	Repo * repository.SinglePostContentRepository
}

func (service * SinglePostContentService) CreateSinglePostContent(singlePostContent *model.SinglePostContent) error {
	err := service.Repo.CreateSinglePostContent(singlePostContent)
	if err != nil {
		return err
	}
	return nil
}


func (service *SinglePostContentService) FindAllContentsForPosts(allPosts []postsModel.Post) []model.SinglePostContent {
	posts := service.Repo.FindAllContentsForPosts(allPosts)
	if posts != nil {
		return posts
	}
	return nil
}
