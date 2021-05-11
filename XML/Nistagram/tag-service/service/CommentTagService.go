package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type CommentTagService struct {
	Repo * repository.CommentTagRepository
}

func (service * CommentTagService) CreateCommentTag(commentTag *model.CommentTag) error {
	err := service.Repo.CreateCommentTag(commentTag)
	if err != nil {
		return err
	}
	return nil
}
