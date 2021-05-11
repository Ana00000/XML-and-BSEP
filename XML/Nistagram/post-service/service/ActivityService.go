package service

import (
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