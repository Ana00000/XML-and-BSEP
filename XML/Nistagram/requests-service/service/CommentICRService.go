package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
)

type CommentICRService struct {
	Repo *repository.CommentICRRepository
}

func (service *CommentICRService) CreateCommentICR(commentICR *model.CommentICR) error {
	err := service.Repo.CreateCommentICR(commentICR)
	if err != nil {
		return err
	}
	return nil
}
