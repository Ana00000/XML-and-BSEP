package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type PostAlbumTagPostAlbumsRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumTagPostAlbumsRepository) CreatePostAlbumTagPostAlbums(postAlbumTagPostAlbums *model.PostAlbumTagPostAlbums) error {
	result := repo.Database.Create(postAlbumTagPostAlbums)
	fmt.Print(result)
	return nil
}