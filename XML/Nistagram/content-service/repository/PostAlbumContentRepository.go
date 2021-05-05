package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostAlbumContentRepository struct {
	Database * gorm.DB
}

func (repo * PostAlbumContentRepository) CreatePostAlbumContent(postAlbumContent *model.PostAlbumContent) error {
	result := repo.Database.Create(postAlbumContent)
	fmt.Print(result)
	return nil
}
