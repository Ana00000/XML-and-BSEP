package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
)

type CommentService struct {
	Repo * repository.CommentRepository
}

func (service * CommentService) CreateComment(comment *model.Comment) error {
	err := service.Repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service * CommentService) FindAllCommentsForPost(postId uuid.UUID) []model.Comment {
	comments := service.Repo.FindAllCommentsForPost(postId)
	return comments
}

func (service * CommentService) FindAllUserComments(userId uuid.UUID) []model.Comment {
	comments := service.Repo.FindAllUserComments(userId)
	return comments
}