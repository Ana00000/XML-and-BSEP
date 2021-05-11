package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
)

type SinglePostService struct {
	Repo * repository.SinglePostRepository
}

func (service * SinglePostService) CreateSinglePost(singlePost *model.SinglePost) error {
	err := service.Repo.CreateSinglePost(singlePost)
	if err != nil {
		return err
	}
	return nil
}
