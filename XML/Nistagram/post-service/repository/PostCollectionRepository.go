package repository

import (
	"fmt"
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
