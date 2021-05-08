package service

import (
	"../model"
	"../repository"
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