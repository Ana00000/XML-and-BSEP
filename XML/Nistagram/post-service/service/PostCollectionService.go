package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
)

type PostCollectionService struct {
	Repo * repository.PostCollectionRepository
}

func (service * PostCollectionService) CreatePostCollection(postCollection *model.PostCollection) error {
	err := service.Repo.CreatePostCollection(postCollection)
	if err != nil {
		return err
	}
	return nil
}

func (service * PostCollectionService) FindAllPostCollectionsForUserRegisteredUser(postCollectionUserId uuid.UUID) []model.PostCollection{
	postCollections := service.Repo.FindAllPostCollectionsForUserRegisteredUser(postCollectionUserId)
	if postCollections != nil {
		return postCollections
	}
	return nil
}