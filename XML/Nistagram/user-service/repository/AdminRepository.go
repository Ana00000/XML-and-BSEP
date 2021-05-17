package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	Database *gorm.DB
}

func (repo *AdminRepository) CreateAdmin(admin *model.Admin) error {
	result := repo.Database.Create(admin)
	fmt.Print(result)
	return nil
}

func (repo *AdminRepository) UpdateAdminProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	result := repo.Database.Model(&model.Admin{}).Where("id = ?", user.ID)
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
	fmt.Println(result.RowsAffected)
	fmt.Println("updating admin profile info")
	return nil
}

func (repo *AdminRepository) UpdateAdminPassword(userId uuid.UUID, salt string, password string) error {
	result := repo.Database.Model(&model.Admin{}).Where("id = ?", userId).Update("salt", salt).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
