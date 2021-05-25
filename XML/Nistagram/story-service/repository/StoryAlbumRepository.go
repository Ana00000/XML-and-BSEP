package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"gorm.io/gorm"
	"time"
)

type StoryAlbumRepository struct {
	Database * gorm.DB
}

func (repo * StoryAlbumRepository) CreateStoryAlbum(storyAlbum *model.StoryAlbum) error {
	result := repo.Database.Create(storyAlbum)
	fmt.Print(result)
	return nil
}

func (repo *StoryAlbumRepository) FindAllAlbumStoriesForUser(userId uuid.UUID) []model.StoryAlbum {
	var storyAlbums []model.StoryAlbum
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&storyAlbums)
	return storyAlbums
}

func (repo *StoryAlbumRepository) FindByID(ID uuid.UUID) *model.StoryAlbum {
	storyAlbum := &model.StoryAlbum{}
	if repo.Database.First(&storyAlbum, "id = ? and is_deleted = ?", ID, false).RowsAffected == 0 {
		return nil
	}
	return storyAlbum
}

func (repo *StoryAlbumRepository) FindAllStoryAlbums() []model.StoryAlbum {
	var storyAlbums []model.StoryAlbum
	repo.Database.Select("*").Find(&storyAlbums)
	return storyAlbums
}

func (repo *StoryAlbumRepository) FindAllPublicAlbumStoriesNotRegisteredUser(allValidUsers []dto.ClassicUserDTO) []model.StoryAlbum {
	var allStoryAlbums = repo.FindAllStoryAlbums()
	var allPublicStoryAlbums []model.StoryAlbum
	var notExpiredStoryAlbums []model.StoryAlbum

	for i:=0;i<len(allStoryAlbums);i++{
		for j:=0; j<len(allValidUsers);j++{
			if allStoryAlbums[i].UserId == allValidUsers[j].ID && allStoryAlbums[i].Type == model.PUBLIC{
				allPublicStoryAlbums = append(allPublicStoryAlbums, allStoryAlbums[i])
			}
		}
	}

	for i:=0; i< len(allPublicStoryAlbums); i++{
		if time.Now().After(allPublicStoryAlbums[i].CreationDate.Add(24 * time.Hour)){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//allPublicStories[i].IsExpired = true
			repo.Database.Model(&model.StoryAlbum{}).Where("id = ?", allPublicStoryAlbums[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", allPublicStoryAlbums[i].ID).Update("is_expired", true)
		} else{
			notExpiredStoryAlbums = append(notExpiredStoryAlbums, allPublicStoryAlbums[i])
		}
	}

	return notExpiredStoryAlbums
}

// FIND ALL NOT DELETED VALID STORY ALBUMS THAT LOGGED IN USER FOLLOWS
func (repo *StoryAlbumRepository) FindAllFollowingStoryAlbums(followings []dto.ClassicUserFollowingsDTO) []model.StoryAlbum {
	var allStoryAlbums = repo.FindAllStoryAlbums()
	var allFollowingStoryAlbums []model.StoryAlbum
	var notExpiredStoryAlbums []model.StoryAlbum

	for i:= 0; i< len(allStoryAlbums); i++{
		for j := 0; j < len(followings); j++{
			if (allStoryAlbums[i].UserId == followings[j].FollowingUserId) && (allStoryAlbums[i].IsDeleted == false){
				allFollowingStoryAlbums = append(allFollowingStoryAlbums, allStoryAlbums[i])
			}
		}
	}

	for i:=0; i< len(allFollowingStoryAlbums); i++{
		if time.Now().After(allFollowingStoryAlbums[i].CreationDate.Add(24 * time.Hour)){
			// PASSED TIME SHOULD SET STORY ALBUM AS EXPIRED
			//allFollowingStoryAlbums[i].IsExpired = true
			repo.Database.Model(&model.StoryAlbum{}).Where("id = ?", allFollowingStoryAlbums[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", allFollowingStoryAlbums[i].ID).Update("is_expired", true)
		} else{
			notExpiredStoryAlbums = append(notExpiredStoryAlbums, allFollowingStoryAlbums[i])
		}
	}
	return notExpiredStoryAlbums
}