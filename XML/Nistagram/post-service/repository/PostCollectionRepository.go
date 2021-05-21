package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type PostCollectionRepository struct {
	Database * gorm.DB
}

func (repo * PostCollectionRepository) CreatePostCollection(postCollection *model.PostCollection) error {
	result := repo.Database.Create(postCollection)
	fmt.Print(result)
	return nil
}

func (repo * PostCollectionRepository) FindAllPostCollectionsForUserRegisteredUser(postCollectionUserId uuid.UUID) []model.PostCollection{
	var postCollections []model.PostCollection
	repo.Database.Select("*").Where("user_id = ?", postCollectionUserId).Find(&postCollections)
	return postCollections
}