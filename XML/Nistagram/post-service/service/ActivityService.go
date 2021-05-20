package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
)

type ActivityService struct {
	Repo * repository.ActivityRepository
}

func (service * ActivityService) CreateActivity(activity *model.Activity) error {
	err := service.Repo.CreateActivity(activity)
	if err != nil {
		return err
	}
	return nil
}

func (service * ActivityService) FindAllLikesForPost(postId uuid.UUID) []model.Activity{
	activities := service.Repo.FindAllLikesForPost(postId)
	if activities != nil {
		return activities
	}
	return nil
}


func (service * ActivityService) FindAllDislikesForPost(postId uuid.UUID) []model.Activity{
	activities := service.Repo.FindAllDislikesForPost(postId)
	if activities != nil {
		return activities
	}
	return nil
}

func (service * ActivityService) FindAllFavoritesForPost(postId uuid.UUID) []model.Activity{
	activities := service.Repo.FindAllFavoritesForPost(postId)
	if activities != nil {
		return activities
	}
	return nil
}

func (service * ActivityService) FindAllActivitiesForPost(postId uuid.UUID) []model.Activity{
	activities := service.Repo.FindAllActivitiesForPost(postId)
	if activities != nil {
		return activities
	}
	return nil
}



