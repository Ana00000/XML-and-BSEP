package repository

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"fmt"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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

func (repo *PostAlbumRepository) FindAllPostAlbums() []model.PostAlbum {
	var postAlbums []model.PostAlbum
	repo.Database.Select("*").Find(&postAlbums)
	return postAlbums
}

func (repo *PostAlbumRepository) FindAllPublicAndFriendsPostAlbumsValid(allValidUsers []userModel.ClassicUser) []model.PostAlbum {
	var allPostAlbums = repo.FindAllPostAlbums()
	var allPublicPostAlbums []model.PostAlbum

	for i:=0;i<len(allPostAlbums);i++{
		for j:=0; j<len(allValidUsers);j++{
			if allPostAlbums[i].UserID == allValidUsers[j].ID {
				allPublicPostAlbums = append(allPublicPostAlbums, allPostAlbums[i])
			}
		}
	}

	return allPublicPostAlbums
}