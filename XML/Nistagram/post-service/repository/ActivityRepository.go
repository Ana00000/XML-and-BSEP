package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	Database * gorm.DB
}

func (repo * ActivityRepository) CreateActivity(activity *model.Activity) error {
	result := repo.Database.Create(activity)
	fmt.Print(result)
	return nil
}

func (repo * ActivityRepository) FindAllLikesForPost(postId uuid.UUID) []model.Activity{
	var likes []model.Activity

	repo.Database.Select("*").Where("post_id = ? and liked_status = ?", postId, 0).Find(&likes)
	return likes
}

func (repo * ActivityRepository) FindAllDislikesForPost(postId uuid.UUID) []model.Activity{
	var dislikes []model.Activity

	repo.Database.Select("*").Where("post_id = ? and liked_status = ?", postId, 1).Find(&dislikes)
	return dislikes
}

func (repo * ActivityRepository) FindAllFavoritesForPost(postId uuid.UUID) []model.Activity{
	var favorites []model.Activity

	repo.Database.Select("*").Where("post_id = ? and is_favorite = ?", postId, true).Find(&favorites)
	return favorites
}