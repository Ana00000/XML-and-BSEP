package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"fmt"
	"gorm.io/gorm"
)

type PostAlbumRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumRepository) CreatePostAlbum(postAlbum *model.PostAlbum) error {
	result := repo.Database.Create(postAlbum)
	fmt.Println(result)
	return nil
}