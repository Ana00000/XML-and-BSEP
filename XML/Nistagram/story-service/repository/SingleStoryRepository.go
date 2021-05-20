package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
	"time"
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


// THIS SHOULD ONLY BE ACCESSED BY THE LOGGED IN USER WHOSE STORY IT IS
// HE CAN SEE HIS EXPIRED STORIES
func (repo *SingleStoryRepository) FindByID(ID uuid.UUID) *model.SingleStory {
	story := &model.SingleStory{}
	if repo.Database.First(&story, "id = ? and is_deleted = ?", ID, false).RowsAffected == 0 {
		return nil
	}

	if story.CreationDate.Add(60 * time.Second).After(time.Now()){
		// PASSED TIME SHOULD SET STORY AS EXPIRED
		//stories[i].IsExpired = true
		repo.Database.Model(&model.SingleStory{}).Where("id = ?", story.ID).Update("is_expired", true)
		repo.Database.Model(&model.Story{}).Where("id = ?", story.ID).Update("is_expired", true)
	}

	return story
}






// DONNEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE



// tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA
func (repo *SingleStoryRepository) FindAllPublicStoriesNotRegisteredUser(allValidUsers []userModel.ClassicUser) []model.SingleStory {
	var allStories = repo.FindAllStories()
	var allPublicStories []model.SingleStory
	var notExpiredStories []model.SingleStory

	for i:=0;i<len(allStories);i++{
		for j:=0; j<len(allValidUsers);j++{
			if allStories[i].UserId == allValidUsers[j].ID && allStories[i].Type == model.PUBLIC{
				allPublicStories = append(allPublicStories, allStories[i])
			}
		}
	}

	for i:=0; i< len(allPublicStories); i++{
		if allPublicStories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//allPublicStories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", allPublicStories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", allPublicStories[i].ID).Update("is_expired", true)
		} else{
			notExpiredStories = append(notExpiredStories, allPublicStories[i])
		}
	}

	return notExpiredStories
}

func (repo *SingleStoryRepository) FindAllStoriesForUserNotReg(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	var notExpiredStories []model.SingleStory
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ? and is_expired = ? and type = ?", userId, false, false, 2).Find(&stories)

	for i:=0; i< len(stories); i++{
		if stories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//stories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
		} else{
			notExpiredStories = append(notExpiredStories, stories[i])
		}
	}

	return notExpiredStories
}


// metoda koja se poziva kada registrovani user udje na profil nekog usera kome je on CLOSE FRIEND
func (repo *SingleStoryRepository) FindAllStoriesForUserCloseFriend(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	var notExpiredStories []model.SingleStory

	repo.Database.Select("*").Where("user_id = ? and is_deleted = ? and is_expired = ?", userId, false, false).Find(&stories)

	for i:=0; i< len(stories); i++{
		if stories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//stories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
		} else{

			notExpiredStories = append(notExpiredStories, stories[i])
		}
	}

	return notExpiredStories
}

// metoda koja se poziva kada registrovani user udje na profil nekog usera koga prati ali nije CLOSE FRIENDS
// ZNACI USLOVI: ILI PUBLIC STORY ILI ALL FRIENDS STORY
func (repo *SingleStoryRepository) FindAllStoriesForUserPublicAllFriends(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	var notExpiredStories []model.SingleStory

	repo.Database.Select("*").Where("user_id = ? and is_deleted = ? and is_expired = ?", userId, false, false).Find(&stories)

	for i:=0; i< len(stories); i++{
		if stories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//stories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
		} else{
			if stories[i].Type == model.PUBLIC || stories[i].Type == model.ALL_FRIENDS{
				notExpiredStories = append(notExpiredStories, stories[i])
			}
		}
	}

	return notExpiredStories
}


func (repo *SingleStoryRepository) FindAllStoriesForUserPublic(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	var notExpiredStories []model.SingleStory

	repo.Database.Select("*").Where("user_id = ? and is_deleted = ? and is_expired = ?", userId, false, false).Find(&stories)

	for i:=0; i< len(stories); i++{
		if stories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//stories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
		} else{
			if stories[i].Type == model.PUBLIC{
				notExpiredStories = append(notExpiredStories, stories[i])
			}
		}
	}

	return notExpiredStories
}

// FIND ALL NOT DELETED VALID STORIES THAT LOGGED IN USER FOLLOWS
func (repo *SingleStoryRepository) FindAllFollowingStories(followings []userModel.ClassicUserFollowings) []model.SingleStory {
	var allStories = repo.FindAllStories()
	var allFollowingStories []model.SingleStory
	var notExpiredStories []model.SingleStory

	for i:= 0; i< len(allStories); i++{
		for j := 0; j < len(followings); j++{
			if (allStories[i].UserId == followings[j].FollowingUserId) && (allStories[i].IsDeleted == false){
				allFollowingStories = append(allFollowingStories, allStories[i])
			}
		}
	}

	for i:=0; i< len(allFollowingStories); i++{
		if allFollowingStories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//allFollowingStories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", allFollowingStories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", allFollowingStories[i].ID).Update("is_expired", true)
		} else{

				notExpiredStories = append(notExpiredStories, allFollowingStories[i])


		}
	}

	return notExpiredStories

}

// FOR MY USER WHEN HE WANTS TO LOOK AT HIS ARCHIVED STORIES
// USED WHEN CLICKING ON A SELECTED USER (YOU CAN SELECT FROM A LIST OF ONLY VALID USERS)
// updates expired stories status but still returns all of them
func (repo *SingleStoryRepository) FindAllStoriesForLoggedUser(userId uuid.UUID) []model.SingleStory {
	var stories []model.SingleStory
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&stories)

	for i:=0; i< len(stories); i++{
		if stories[i].CreationDate.Add(60 * time.Second).After(time.Now()){
			// PASSED TIME SHOULD SET STORY AS EXPIRED
			//stories[i].IsExpired = true
			repo.Database.Model(&model.SingleStory{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
			repo.Database.Model(&model.Story{}).Where("id = ?", stories[i].ID).Update("is_expired", true)
		}
	}

	return stories
}

