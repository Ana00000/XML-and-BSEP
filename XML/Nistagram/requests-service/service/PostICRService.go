package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
)

type PostICRService struct {
	Repo *repository.PostICRRepository
}

func (service *PostICRService) CreatePostICR(postICR *model.PostICR) error {
	err := service.Repo.CreatePostICR(postICR)
	if err != nil {
		return err
	}
	return nil
}
