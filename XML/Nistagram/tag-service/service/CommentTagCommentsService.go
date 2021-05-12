package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type CommentTagCommentsService struct {
	Repo * repository.CommentTagCommentsRepository
}

func (service * CommentTagCommentsService) CreateCommentTagComments(commentTagComments *model.CommentTagComments) error {
	err := service.Repo.CreateCommentTagComments(commentTagComments)
	if err != nil {
		return err
	}
	return nil
}
