package repository

import (
	"fmt"
	"github.com/google/uuid"
	requestModel "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserFollowingsRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserFollowingsRepository) CreateClassicUserFollowings(classicUserFollowings *model.ClassicUserFollowings) error {
	result := repo.Database.Create(classicUserFollowings)
	fmt.Print(result)
	return nil
}
func (repo *ClassicUserFollowingsRepository) FindById(id string) *model.ClassicUserFollowings {
	following := &model.ClassicUserFollowings{}
	if repo.Database.First(&following, "id = ?", id).RowsAffected == 0{
		return nil
	}
	return following
}

func (repo * ClassicUserFollowingsRepository) FindAllFollowingsForUser(userId uuid.UUID) []model.ClassicUserFollowings{
	var followings []model.ClassicUserFollowings
	repo.Database.Select("*").Where("classic_user_id = ?", userId).Find(&followings)
	return followings
}

func (repo * ClassicUserFollowingsRepository) FindAllFollowingsPerUser(userId uuid.UUID) []model.ClassicUserFollowings{
	var followings []model.ClassicUserFollowings
	repo.Database.Select("*").Where("following_user_id = ?", userId).Find(&followings)
	return followings
}

func (repo * ClassicUserFollowingsRepository) CheckFollowingStatus(classicUserId uuid.UUID, followingUserId uuid.UUID, followRequests []requestModel.FollowRequest) string{
	var allFollowingForUser = repo.FindAllFollowingsForUser(classicUserId)

	fmt.Println("DUZINA")
	fmt.Println(len(followRequests))
	following, doneFollowing := repo.CheckIfFollowing(allFollowingForUser, followingUserId)
	if doneFollowing {
		return following
	}

	pending, donePending := repo.CheckIfPending(followRequests, followingUserId)
	if donePending {
		return pending
	}

	fmt.Println("NISTA")
	return "NOT FOLLOWING"
}

func (repo *ClassicUserFollowingsRepository) CheckIfPending(followRequests []requestModel.FollowRequest, followingUserId uuid.UUID) (string, bool) {
	for i := 0; i < len(followRequests); i++ {
		if followRequests[i].FollowerUserId == followingUserId && followRequests[i].FollowRequestStatus == requestModel.PENDING {
			fmt.Println("PENDING JE")
			return "PENDING", true
		}
	}
	return "", false
}

func (repo *ClassicUserFollowingsRepository) CheckIfFollowing(allFollowingForUser []model.ClassicUserFollowings, followingUserId uuid.UUID) (string, bool) {
	for i := 0; i < len(allFollowingForUser); i++ {
		if allFollowingForUser[i].FollowingUserId == followingUserId {
			fmt.Println("FOLLOWING JE")
			return "FOLLOWING", true
		}
	}
	return "", false
}

func (repo * ClassicUserFollowingsRepository) FindAllValidFollowingsForUser(userId uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUserFollowings{
	var followings  = repo.FindAllFollowingsForUser(userId)
	var validFollowings []model.ClassicUserFollowings
	for i := 0; i < len(followings); i++ {
		fmt.Println("FindAllValidFollowingsForUser followings "+followings[i].ClassicUserId.String())
	}
	for i := 0; i < len(allValidUsers); i++ {
		fmt.Println("FindAllValidFollowingsForUser allValidUsers "+allValidUsers[i].ID.String())
	}
	for i := 0; i < len(allValidUsers); i++{
		for j :=0; j < len(followings); j++{

			if allValidUsers[i].ID == followings[j].FollowingUserId{
				fmt.Println("Unutar if-a "+allValidUsers[i].Username)
				validFollowings = append(validFollowings, followings[j])
			}
		}
	}

	return validFollowings

}

func (repo *ClassicUserFollowingsRepository) CheckIfFollowingPostStory(followingUserId uuid.UUID, classicUserId uuid.UUID) bool{
	follower := &model.ClassicUserFollowings{}
	if repo.Database.First(&follower, "following_user_id = ? and classic_user_id = ?", followingUserId, classicUserId).RowsAffected == 0{
		return false
	}
	return true
}

func (repo * ClassicUserFollowingsRepository) FindAllUserWhoFollow(userId uuid.UUID) []model.ClassicUserFollowings{
	var followings []model.ClassicUserFollowings
	repo.Database.Select("*").Where("classic_user_id = ?", userId).Find(&followings)
	return followings
}

func (repo * ClassicUserFollowingsRepository) FindAllUserWhoFollowUserId(userId uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUserFollowings{
	var followings  = repo.FindAllUserWhoFollow(userId)
	var validFollowings []model.ClassicUserFollowings

	for i := 0; i < len(allValidUsers); i++{
		classicUserFollowings, exist := ExsistInList(allValidUsers[i], followings)
		if exist {
			fmt.Println("Dodaje u listu")
			validFollowings = append(validFollowings,classicUserFollowings)
		}
	}

	return validFollowings

}

func ExsistInList(user model.ClassicUser,followings []model.ClassicUserFollowings) (model.ClassicUserFollowings, bool){
	var classicUserFollowings model.ClassicUserFollowings
	for i := 0; i < len(followings); i++ {

		if user.ID == followings[i].FollowingUserId{
			fmt.Println(followings[i].FollowingUserId.String()+" ExsistInList")
			return followings[i], true
		}
	}
	return classicUserFollowings, false
}
