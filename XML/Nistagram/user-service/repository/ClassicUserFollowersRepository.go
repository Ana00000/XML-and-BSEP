package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserFollowersRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserFollowersRepository) CreateClassicUserFollowers(classicUserFollowers *model.ClassicUserFollowers) error {
	result := repo.Database.Create(classicUserFollowers)
	fmt.Print(result)
	return nil
}

func (repo *ClassicUserFollowersRepository) FindById(id string) *model.ClassicUserFollowers {
	follower := &model.ClassicUserFollowers{}
	if repo.Database.First(&follower, "id = ?", id).RowsAffected == 0{
		return nil
	}
	return follower
}

func (repo * ClassicUserFollowersRepository) FindAllFollowersForUser(userId uuid.UUID) []model.ClassicUserFollowers{
	var followers []model.ClassicUserFollowers
	repo.Database.Select("*").Where("classic_user_id = ?", userId).Find(&followers)
	return followers
}

//CheckIfFollowers
func (repo *ClassicUserFollowersRepository) CheckIfFollowers(classicUserId uuid.UUID, followerUserId uuid.UUID) bool {
	follower := &model.ClassicUserFollowers{}
	if repo.Database.First(&follower, "classic_user_id = ? and follower_user_id = ?", classicUserId, followerUserId).RowsAffected == 0{
		return false
	}
	return true
}

