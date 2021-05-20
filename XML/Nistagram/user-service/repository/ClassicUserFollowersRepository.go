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

//metoda koja dobavlja sve potencijalne close friends (mutual follow izmedju usera)
// MUTUAL FOLLOWERS FOR FIRST USER
func (repo *ClassicUserFollowersRepository) FindAllMutualFollowerForUser(userId uuid.UUID) []model.ClassicUserFollowers {

	var allFollowersForFirstUser = repo.FindAllFollowersForUser(userId)

	fmt.Println(allFollowersForFirstUser)
	var mutualFollowers []model.ClassicUserFollowers

	for i:=0; i<len(allFollowersForFirstUser); i++{
		var allFollowersForSecondUser = repo.FindAllFollowersForUser(allFollowersForFirstUser[i].FollowerUserId)
fmt.Println(allFollowersForSecondUser)
		for j:=0; j<len(allFollowersForSecondUser);j++{
			if allFollowersForSecondUser[j].FollowerUserId == userId{
				mutualFollowers = append(mutualFollowers, allFollowersForFirstUser[i])
			}
		}
	}

	return mutualFollowers
}



