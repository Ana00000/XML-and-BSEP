package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
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


func (service *SinglePostContentService) FindAllContentsForPosts(allPosts []dto.SinglePostDTO) []model.SinglePostContent {
	posts := service.Repo.FindAllContentsForPosts(allPosts)
	if posts != nil {
		return posts
	}
	return nil
}

func (service *SinglePostContentService) FindAllContentsForPost(post *dto.SinglePostDTO) []model.SinglePostContent {
	posts := service.Repo.FindAllContentsForPost(post)
	if posts != nil {
		return posts
	}
	return nil
}

