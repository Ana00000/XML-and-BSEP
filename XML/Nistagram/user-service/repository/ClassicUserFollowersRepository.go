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

func (repo * ClassicUserFollowersRepository) FindAllFollowersInfoForUser(userId uuid.UUID) []model.User{
	var followers = repo.FindAllFollowersForUser(userId)
	var users []model.User
	repo.Database.Select("*").Find(&users)
	var followerUser []model.User
	
	for i := 0; i < len(users); i++{
		for j :=0; j<len(followers); j++{
			if users[i].ID == followers[j].ClassicUserId{
				followerUser = append(followerUser, users[i])
			}

		}
	}
	return followerUser
}
