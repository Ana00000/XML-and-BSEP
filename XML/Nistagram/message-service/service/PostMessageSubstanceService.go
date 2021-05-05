package service

import (
	"../model"
	"../repository"
)

type PostMessageSubstanceService struct {
	Repo *repository.PostMessageSubstanceRepository
}

func (service * PostMessageSubstanceService) CreatePostMessageSubstance(postMessageSubstance *model.PostMessageSubstance) error {
	service.Repo.CreatePostMessageSubstance(postMessageSubstance)
	return nil
}


