package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type PostAlbumTagRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumTagRepository) CreatePostAlbumTag(postAlbumTag *model.PostAlbumTag) error {
	result := repo.Database.Create(postAlbumTag)
	fmt.Print(result)
	return nil
}