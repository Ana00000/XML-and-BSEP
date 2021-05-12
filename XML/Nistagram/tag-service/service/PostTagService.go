package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type PostTagService struct {
	Repo * repository.PostTagRepository
}

func (service * PostTagService) CreatePostTag(postTag *model.PostTag) error {
	err := service.Repo.CreatePostTag(postTag)
	if err != nil {
		return err
	}
	return nil
}
