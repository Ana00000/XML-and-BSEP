package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type SingleStoryRepository struct {
	Database * gorm.DB
}

func (repo * SingleStoryRepository) CreateSingleStory(singleStory *model.SingleStory) error {
	result := repo.Database.Create(singleStory)
	fmt.Print(result)
	return nil
}

func (repo *SingleStoryRepository) FindAllStories() []model.SingleStory {
	var stories []model.SingleStory
	repo.Database.Select("*").Find(&stories)
	return stories
}

func (repo *SingleStoryRepository) FindByID(ID uuid.UUID) *model.SingleStory {
	story := &model.SingleStory{}
	if repo.Database.First(&story, "id = ? and is_deleted = ?", ID, false).RowsAffected == 0 {
		return nil
	}
	return story
}

// USED WHEN CLICKING ON A SELECTED USER (YOU CAN SELECT FROM A LIST OF ONLY VALID USERS)
func (repo *SingleStoryRepository) FindAllStoriesForUser(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&stories)
	return stories
}


// FIND ALL NOT DELETED VALID STORIES THAT LOGGED IN USER FOLLOWS
func (repo *SingleStoryRepository) FindAllFollowingStories(followings []userModel.ClassicUserFollowings) []model.SingleStory {
	var allStories = repo.FindAllStories()
	var allFollowingStories []model.SingleStory

	for i:= 0; i< len(allStories); i++{
		for j := 0; j < len(followings); j++{
			if (allStories[i].UserId == followings[j].FollowingUserId) && (allStories[i].IsDeleted == false){
				allFollowingStories = append(allFollowingStories, allStories[i])
			}
		}
	}
	return allFollowingStories
}

func (repo *SingleStoryRepository) FindAllPublicStoriesNotRegisteredUser(allValidUsers []userModel.ClassicUser) []model.SingleStory {
	var allStories = repo.FindAllStories()
	var allPublicStories []model.SingleStory

	for i:=0;i<len(allStories);i++{
		for j:=0; j<len(allValidUsers);j++{
			if allStories[i].UserId == allValidUsers[j].ID {
				allPublicStories = append(allPublicStories, allStories[i])
			}
		}
	}

	return allPublicStories
}