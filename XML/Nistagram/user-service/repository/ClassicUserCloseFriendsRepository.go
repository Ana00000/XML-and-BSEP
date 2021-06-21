package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserCloseFriendsRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserCloseFriendsRepository) CreateClassicUserCloseFriends(classicUserCloseFriends *model.ClassicUserCloseFriends) error {
	result := repo.Database.Create(classicUserCloseFriends)
	fmt.Print(result)
	return nil
}

// FindAllCloseFriendsForUser
func (repo * ClassicUserCloseFriendsRepository) FindAllCloseFriendsForUser(userId uuid.UUID) []model.ClassicUserCloseFriends{
	var closeFriends []model.ClassicUserCloseFriends
	repo.Database.Select("*").Where("classic_user_id = ?", userId).Find(&closeFriends)
	return closeFriends
}


//CheckIfCloseFriend
func (repo * ClassicUserCloseFriendsRepository) CheckIfCloseFriend(classicUserId uuid.UUID, closeFriendUserId uuid.UUID) bool{
	var closeFriends *model.ClassicUserCloseFriends

	if repo.Database.First(&closeFriends, "classic_user_id = ? and close_friend_user_id = ?", classicUserId, closeFriendUserId).RowsAffected == 0{
		return false
	}
	return true

}

func (repo ClassicUserCloseFriendsRepository) FindCloseFriendByUsersIDs(closeFriendID uuid.UUID,classicUserID uuid.UUID) *model.ClassicUserCloseFriends {
	classicUserCloseFriends := &model.ClassicUserCloseFriends{}
	if repo.Database.First(&classicUserCloseFriends, "close_friend_user_id = ? and classic_user_id = ?", closeFriendID, classicUserID).RowsAffected == 0{
		return nil
	}
	return classicUserCloseFriends
}

func (repo *ClassicUserCloseFriendsRepository) RemoveClassicUserCloseFriend(id uuid.UUID) {
	repo.Database.Delete(&model.ClassicUserCloseFriends{}, id)
}


