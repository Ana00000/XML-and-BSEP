package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
	"time"
)

type AdminRepository struct {
	Database * gorm.DB
}

func (repo * AdminRepository) CreateAdmin(admin *model.Admin) error {
	result := repo.Database.Create(admin)
	fmt.Print(result)
	return nil
}

func (repo * AdminRepository) UpdateAdminProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, user.DateOfBirth)

	result := repo.Database.Model(&model.Admin{}).Where("id = ?", user.ID).Update("username", user.Username).Update("phoneNumber", user.PhoneNumber).Update("firstName", user.FirstName).Update("lastName", user.LastName).Update("gender", gender).Update("dateOfBirth", dateOfBirth).Update("website", user.Website).Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating admin profile info")
}

func (repo *AdminRepository) UpdateAdminPassword(userId uuid.UUID, password string) error {
	result := repo.Database.Model(&model.Admin{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}