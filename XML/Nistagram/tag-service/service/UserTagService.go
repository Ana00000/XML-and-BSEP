package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type UserTagService struct {
	Repo *repository.UserTagRepository
}

func (service *UserTagService) CreateUserTag(userTag *model.UserTag) error {
	if err := service.Repo.CreateUserTag(userTag); err != nil {
		return err
	}
	return nil
}