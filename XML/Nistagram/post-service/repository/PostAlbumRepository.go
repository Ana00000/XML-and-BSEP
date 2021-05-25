package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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

func (repo *PostAlbumRepository) FindAllPublicAndFriendsPostAlbumsValid(allValidUsers []dto.ClassicUserDTO) []model.PostAlbum {
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

// FIND ALL NOT DELETED VALID POST ALBUMS THAT LOGGED IN USER FOLLOWS
func (repo *PostAlbumRepository) FindAllFollowingPostAlbums(followings []dto.ClassicUserFollowingsFullDTO) []model.PostAlbum {
	var allPostAlbums = repo.FindAllPostAlbums()
	var allFollowingPostAlbums []model.PostAlbum

	for i:= 0; i< len(allPostAlbums); i++{
		for j := 0; j < len(followings); j++{
			if (allPostAlbums[i].UserID == followings[j].FollowingUserId) && (allPostAlbums[i].IsDeleted == false){
				allFollowingPostAlbums = append(allFollowingPostAlbums, allPostAlbums[i])
			}
		}
	}
	return allFollowingPostAlbums
}