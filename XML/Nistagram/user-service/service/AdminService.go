package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
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

func (service *AdminService) UpdateAdminProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateAdminProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}
func (service *AdminService) UpdateAdminPassword(userId uuid.UUID, password string) error {
	err := service.Repo.UpdateAdminPassword(userId,password)
	if err != nil {
		return err
	}
	return nil
}