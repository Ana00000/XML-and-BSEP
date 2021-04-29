package service

import (
	"../model"
	"../repository"
)

type AdminService struct {
	Repo * repository.AdminRepository
}

func (service * AdminService) CreateAdmin(admin *model.Admin) error {
	service.Repo.CreateAdmin(admin)
	return nil
}