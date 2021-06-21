package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
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

func (repo * ClassicUserFollowingsRepository) CheckFollowingStatus(classicUserId uuid.UUID, followingUserId uuid.UUID, followRequests []dto.FollowRequestForUserDTO) string{
	var allFollowingForUser = repo.FindAllFollowingsForUser(classicUserId)

	fmt.Println("DUZINA u CheckFollowingStatus" + string(len(followRequests)))
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

func (repo *ClassicUserFollowingsRepository) CheckIfPending(followRequests []dto.FollowRequestForUserDTO, followingUserId uuid.UUID) (string, bool) {
	for i := 0; i < len(followRequests); i++ {
		fmt.Println(followRequests[i].FollowRequestStatus)
		if followRequests[i].FollowerUserId == followingUserId && followRequests[i].FollowRequestStatus == "PENDING" {
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

func (repo * ClassicUserFollowingsRepository) FindAllUsersWhoFollowUserId(userId uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUser{
	var followings  = repo.FindAllUserWhoFollow(userId)
	var validFollowingsClassicUsers []model.ClassicUser

	for i := 0; i < len(allValidUsers); i++{
		_, exist := ExsistInList(allValidUsers[i], followings)
		if exist {
			fmt.Println("Dodaje u listu")
			validFollowingsClassicUsers = append(validFollowingsClassicUsers,allValidUsers[i])
		}
	}

	return validFollowingsClassicUsers
}

func (repo *ClassicUserFollowingsRepository) FindFollowingByUsersIDs(followingUserID uuid.UUID,classicUserID uuid.UUID) *model.ClassicUserFollowings {
	classicUserFollowings := &model.ClassicUserFollowings{}
	if repo.Database.First(&classicUserFollowings, "following_user_id = ? and classic_user_id = ?", followingUserID, classicUserID).RowsAffected == 0{
		return nil
	}
	return classicUserFollowings
}

func (repo *ClassicUserFollowingsRepository) RemoveClassicUserFollowing(id uuid.UUID) {
	repo.Database.Delete(&model.ClassicUserFollowings{}, id)
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
