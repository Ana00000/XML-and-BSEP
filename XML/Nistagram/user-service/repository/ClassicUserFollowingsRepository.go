package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (repo * ClassicUserFollowingsRepository) FindAllFollowersForUser(userId uuid.UUID) []model.ClassicUserFollowings{
	var followings []model.ClassicUserFollowings
	repo.Database.Select("id = ?", userId).Find(&followings)
	return followings
}