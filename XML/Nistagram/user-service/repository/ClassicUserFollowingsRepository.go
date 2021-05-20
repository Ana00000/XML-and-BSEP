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
