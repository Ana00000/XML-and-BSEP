package repository

import (
	"github.com/google/uuid"
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

func (repo *PostAlbumRepository) FindAllAlbumPostsForUser(userId uuid.UUID) []model.PostAlbum {
	var postAlbums []model.PostAlbum
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&postAlbums)
	return postAlbums
}

func (repo *PostAlbumRepository) FindByID(ID uuid.UUID) *model.PostAlbum {
	postAlbum := &model.PostAlbum{}
	if repo.Database.First(&postAlbum, "id = ? and is_deleted = ?", ID, false).RowsAffected == 0 {
		return nil
	}
	return postAlbum
}
