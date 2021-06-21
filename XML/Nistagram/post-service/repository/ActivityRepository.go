package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	Database *gorm.DB
}

func (repo *ActivityRepository) CreateActivity(activity *model.Activity) error {
	result := repo.Database.Create(activity)
	fmt.Print(result)
	return nil
}

func (repo *ActivityRepository) FindAllLikesForPost(postId uuid.UUID) []model.Activity {
	var likes []model.Activity

	repo.Database.Select("*").Where("post_id = ? and liked_status = ?", postId, 0).Find(&likes)
	return likes
}

func (repo *ActivityRepository) FindAllDislikesForPost(postId uuid.UUID) []model.Activity {
	var dislikes []model.Activity

	repo.Database.Select("*").Where("post_id = ? and liked_status = ?", postId, 1).Find(&dislikes)
	return dislikes
}

func (repo *ActivityRepository) FindAllFavoritesForPost(postId uuid.UUID) []model.Activity {
	var favorites []model.Activity

	repo.Database.Select("*").Where("post_id = ? and is_favorite = ?", postId, true).Find(&favorites)
	return favorites
}

func (repo *ActivityRepository) FindAllActivitiesForPost(postId uuid.UUID) []model.Activity {
	var allReactions []model.Activity
	repo.Database.Select("*").Where("post_id = ?", postId).Find(&allReactions)
	return allReactions
}

func (repo *ActivityRepository) FindAllLikedPostsByUserId(userId uuid.UUID) []model.Activity {
	var allLikedPostActivities []model.Activity
	repo.Database.Select("*").Where("user_id = ? and liked_status = ?", userId, model.LIKED).Find(&allLikedPostActivities)
	return allLikedPostActivities
}

func (repo *ActivityRepository) FindAllDislikedPostsByUserId(userId uuid.UUID) []model.Activity {
	var allDislikedPostActivities []model.Activity
	repo.Database.Select("*").Where("user_id = ? and liked_status = ?", userId, model.DISLIKED).Find(&allDislikedPostActivities)
	return allDislikedPostActivities
}

func (repo *ActivityRepository) UpdateActivity(activity *dto.ActivityDTO) error {
	result := repo.Database.Model(&model.Activity{}).Where("id = ?", activity.ID).Update("liked_status", activity.LikedStatus).Update("is_favorite", activity.IsFavorite)
	fmt.Print(result)
	return nil
}
