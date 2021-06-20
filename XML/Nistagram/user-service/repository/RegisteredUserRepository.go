package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type RegisteredUserRepository struct {
	Database * gorm.DB
}

func (repo * RegisteredUserRepository) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	result := repo.Database.Create(registeredUser)
	fmt.Print(result)
	return nil
}

func (repo *RegisteredUserRepository) UpdateRegisteredUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.RegisteredUser{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo * RegisteredUserRepository) UpdateRegisteredUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	result := repo.Database.Model(&model.RegisteredUser{}).Where("id = ?", user.ID)
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

func (repo *RegisteredUserRepository) UpdateRegisteredUserPassword(userId uuid.UUID, salt string, password string) error {
	result := repo.Database.Model(&model.RegisteredUser{}).Where("id = ?", userId).Update("salt", salt).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *RegisteredUserRepository) UpdateUserCategory(userId uuid.UUID, category model.RegisteredUserCategory) error {
	result := repo.Database.Model(&model.RegisteredUser{}).Where("id = ?", userId).Update("registered_user_category", category)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *RegisteredUserRepository) UpdateOfficialDocumentPath(id uuid.UUID, path string) error {
	result := repo.Database.Model(&model.RegisteredUser{}).Where("id = ?", id).Update("official_document_path", path)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}