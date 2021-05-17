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

func (repo * ClassicUserFollowingsRepository) CheckIfFollowingUser(classicUserId uuid.UUID, followingUserId uuid.UUID, followRequests []requestModel.FollowRequest) string{
	var allFollowingForUser = repo.FindAllFollowingsForUser(classicUserId)

	for i := 0; i < len(allFollowingForUser); i++{
		if allFollowingForUser[i].FollowingUserId == followingUserId{
			return "FOLLOWING"
		}
	}

	for i:=0; i< len(followRequests); i++{
		if followRequests[i].FollowerUserId == followingUserId && followRequests[i].FollowRequestStatus == requestModel.PENDING{
			return "PENDING"
		}
	}

	return "NOT FOLLOWING"
}
