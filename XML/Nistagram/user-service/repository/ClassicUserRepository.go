package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
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

	result := repo.Database.Model(&model.ClassicUser{}).Where("id = ?", user.ID)
	result.Update("username", user.Username)

	fmt.Println(result.RowsAffected)
	result.Update("phone_number", user.PhoneNumber)
	fmt.Println(result.RowsAffected)
	result.Update("first_name", user.FirstName)
	fmt.Println(result.RowsAffected)
	result.Update("last_name", user.LastName)
	fmt.Println(result.RowsAffected)
	result.Update("gender", gender)
	fmt.Println(result.RowsAffected)
	result.Update("date_of_birth", user.DateOfBirth)
	fmt.Println(result.RowsAffected)
	result.Update("website", user.Website)
	fmt.Println(result.RowsAffected)
	result.Update("biography", user.Biography)
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
