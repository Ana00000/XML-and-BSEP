package service

import (
	"../model"
	"../repository"
)

type ActivityService struct {
	Repo * repository.ActivityRepository
}

func (service * ActivityService) CreateActivity(activity *model.Activity) error {
	service.Repo.CreateActivity(activity)
	return nil
}