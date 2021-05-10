package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type AdminService struct {
	Repo * repository.AdminRepository
}

func (service * AdminService) CreateAdmin(admin *model.Admin) error {
	err := service.Repo.CreateAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}