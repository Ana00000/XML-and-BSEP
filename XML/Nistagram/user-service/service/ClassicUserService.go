package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserService struct {
	Repo *repository.ClassicUserRepository
}

func (service *ClassicUserService) CreateClassicUser(classicUser *model.ClassicUser) error {
	err := service.Repo.CreateClassicUser(classicUser)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateClassicUserConfirmed(userId, isConfirmed)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateClassicUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserPassword(userId uuid.UUID, salt string, password string) error {
	err := service.Repo.UpdateClassicUserPassword(userId, salt, password)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) FindSelectedUserById(id uuid.UUID) *dto.SelectedUserDTO {
	user := service.Repo.FindSelectedUserById(id)
	return user
}

func (service *ClassicUserService) FindAllUsersButLoggedIn(userId uuid.UUID) []model.ClassicUser {
	users := service.Repo.FindAllUsersButLoggedIn(userId)
	if users != nil {
		return users
	}
	return nil
}

func (service *ClassicUserService) FindClassicUserByUserName(userName string) *model.ClassicUser {
	user := service.Repo.FindClassicUserByUserName(userName)
	return user
}

func (service *ClassicUserService) CheckIfUserValid(userId uuid.UUID) bool {
	checkIfValid := service.Repo.CheckIfUserValid(userId)
	return checkIfValid
}

func (service *ClassicUserService) FinAllValidUsers() []model.ClassicUser {
	users := service.Repo.FindAllValidUsers()
	if users != nil {
		return users
	}
	return nil
}

func (service *ClassicUserService) FindAllUsersByFollowingIds(userIds []model.ClassicUserFollowings) []model.ClassicUser {
	users := service.Repo.FindAllUsersByFollowingIds(userIds)
	if users != nil {
		fmt.Println("Pronadjeni korisnici")
		return users
	}
	return nil
}

func (service *ClassicUserService) FindById(userId uuid.UUID) *model.ClassicUser {
	user := service.Repo.FindById(userId)
	return user
}


