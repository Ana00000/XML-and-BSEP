package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
	"time"
)

type ClassicUserRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserRepository) CreateClassicUser(classicUser *model.ClassicUser) error {
	result := repo.Database.Create(classicUser)
	fmt.Print(result)
	return nil
}


func (repo *ClassicUserRepository) UpdateClassicUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo * ClassicUserRepository) UpdateClassicUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, user.DateOfBirth)

	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ?", user.ID).Update("username", user.Username).Update("phoneNumber", user.PhoneNumber).Update("firstName", user.FirstName).Update("lastName", user.LastName).Update("gender", gender).Update("dateOfBirth", dateOfBirth).Update("website", user.Website).Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating profile info")
	return nil
}

func (repo *ClassicUserRepository) UpdateClassicUserPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
